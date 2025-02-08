package keeper

import (
	"context"
	"fmt"

	"beep/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MintTokens(goCtx context.Context, msg *types.MsgMintTokens) (*types.MsgMintTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "denom does not exist")
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if valFound.Supply+msg.Amount > valFound.MaxSupply {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Cannot mint more than Max Supply")
	}
	moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)

	recipientAddress, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	var mintCoins sdk.Coins

	mintCoins = mintCoins.Add(sdk.NewCoin(msg.Denom, math.NewInt(int64(msg.Amount))))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return nil, err
	}

	// Emit event for token mint
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMintTokens,
			sdk.NewAttribute(types.AttributeKeyDenomOwner, fmt.Sprint(msg.Owner)),
			sdk.NewAttribute(types.AttributeKeyDenomName, fmt.Sprint(msg.Denom)),
		),
	)

	if err := k.bankKeeper.SendCoins(ctx, moduleAcct, recipientAddress, mintCoins); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTransferTokens,
			sdk.NewAttribute(types.AttributeKeySender, fmt.Sprint(moduleAcct)),
			sdk.NewAttribute(types.AttributeKeyReceiver, fmt.Sprint(msg.Recipient)),
		),
	)

	var denom = types.Denom{
		Owner:              valFound.Owner,
		Denom:              valFound.Denom,
		Description:        valFound.Description,
		MaxSupply:          valFound.MaxSupply,
		Supply:             valFound.Supply + msg.Amount,
		Precision:          valFound.Precision,
		Ticker:             valFound.Ticker,
		Url:                valFound.Url,
		CanChangeMaxSupply: valFound.CanChangeMaxSupply,
	}

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgMintTokensResponse{}, nil

}
