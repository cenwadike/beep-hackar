// package keeper

// import (
// 	"context"

// 	"beep/x/tokenfactory/types"

// 	errorsmod "cosmossdk.io/errors"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
// )

// func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)

// 	// Check if the value already exists
// 	_, isFound := k.GetDenom(
// 		ctx,
// 		msg.Denom,
// 	)
// 	if isFound {
// 		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Denom already exist")
// 	}

// 	var denom = types.Denom{
// 		Owner:              msg.Owner,
// 		Denom:              msg.Denom,
// 		Description:        msg.Description,
// 		Ticker:             msg.Ticker,
// 		Precision:          msg.Precision,
// 		Url:                msg.Url,
// 		MaxSupply:          msg.MaxSupply,
// 		Supply:             0,
// 		CanChangeMaxSupply: msg.CanChangeMaxSupply,
// 	}

// 	k.SetDenom(
// 		ctx,
// 		denom,
// 	)
// 	return &types.MsgCreateDenomResponse{}, nil
// }

// func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
// 	ctx := sdk.UnwrapSDKContext(goCtx)

// 	// Check if the value exists
// 	valFound, isFound := k.GetDenom(
// 		ctx,
// 		msg.Denom,
// 	)
// 	if !isFound {
// 		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "Denom to update not found")
// 	}

// 	// Checks if the msg owner is the same as the current owner
// 	if msg.Owner != valFound.Owner {
// 		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
// 	}

// 	if !valFound.CanChangeMaxSupply && valFound.MaxSupply != msg.MaxSupply {
// 		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "cannot change maxsupply")
// 	}
// 	if !valFound.CanChangeMaxSupply && msg.CanChangeMaxSupply {
// 		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot revert change maxsupply flag")
// 	}

// 	var denom = types.Denom{
// 		Owner:              msg.Owner,
// 		Denom:              msg.Denom,
// 		Description:        msg.Description,
// 		Ticker:             valFound.Ticker,
// 		Precision:          valFound.Precision,
// 		Url:                msg.Url,
// 		MaxSupply:          msg.MaxSupply,
// 		Supply:             valFound.Supply,
// 		CanChangeMaxSupply: msg.CanChangeMaxSupply,
// 	}

// 	k.SetDenom(ctx, denom)

// 	return &types.MsgUpdateDenomResponse{}, nil
// }

package keeper

import (
	"context"
	"fmt"

	"beep/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Denom already exist")
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Ticker:             msg.Ticker,
		Precision:          msg.Precision,
		Url:                msg.Url,
		MaxSupply:          msg.MaxSupply,
		Supply:             0,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	k.SetDenom(
		ctx,
		denom,
	)

	// Emit event that new denom is created
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateDenom,
			sdk.NewAttribute(types.AttributeKeyDenomOwner, fmt.Sprint(msg.Owner)),
			sdk.NewAttribute(types.AttributeKeyDenomName, fmt.Sprint(msg.Denom)),
			sdk.NewAttribute(types.AttributeKeyDenomDescription, fmt.Sprint(msg.Description)),
			sdk.NewAttribute(types.AttributeKeyDenomTicker, fmt.Sprint(msg.Ticker)),
			sdk.NewAttribute(types.AttributeKeyDenomPrecision, fmt.Sprint(msg.Precision)),
			sdk.NewAttribute(types.AttributeKeyDenomMaxSupply, fmt.Sprint(msg.MaxSupply)),
			sdk.NewAttribute(types.AttributeKeyDenomCanChangeMaxSupply, fmt.Sprint(msg.CanChangeMaxSupply)),
		),
	)

	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "Denom to update not found")
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if !valFound.CanChangeMaxSupply && valFound.MaxSupply != msg.MaxSupply {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "cannot change maxsupply")
	}
	if !valFound.CanChangeMaxSupply && msg.CanChangeMaxSupply {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot revert change maxsupply flag")
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Ticker:             valFound.Ticker,
		Precision:          valFound.Precision,
		Url:                msg.Url,
		MaxSupply:          msg.MaxSupply,
		Supply:             valFound.Supply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	k.SetDenom(ctx, denom)

	// Emit event for denom update
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateDenom,
			sdk.NewAttribute(types.AttributeKeyDenomOwner, fmt.Sprint(msg.Owner)),
			sdk.NewAttribute(types.AttributeKeyDenomName, fmt.Sprint(msg.Denom)),
			sdk.NewAttribute(types.AttributeKeyDenomDescription, fmt.Sprint(msg.Description)),
			sdk.NewAttribute(types.AttributeKeyDenomMaxSupply, fmt.Sprint(msg.MaxSupply)),
			sdk.NewAttribute(types.AttributeKeyDenomCanChangeMaxSupply, fmt.Sprint(msg.CanChangeMaxSupply)),
		),
	)

	return &types.MsgUpdateDenomResponse{}, nil
}
