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

func (k msgServer) BurnTokens(goCtx context.Context, msg *types.MsgBurnTokens) (*types.MsgBurnTokensResponse, error) {
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
	if msg.Signer != valFound.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	burnerAddress, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, err
	}

	var mintCoins sdk.Coins
	mintCoins = mintCoins.Add(sdk.NewCoin(msg.Denom, math.NewInt(int64(msg.Amount))))

	moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName)

	if err := k.bankKeeper.SendCoins(ctx, burnerAddress, moduleAcct, mintCoins); err != nil {
		return nil, err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, mintCoins)
	if err != nil {
		return nil, err
	}

	// Emit event for token burn
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeBurnTokens,
			sdk.NewAttribute(types.AttributeKeyBurner, fmt.Sprint(msg.Signer)),
			sdk.NewAttribute(types.AttributeKeyDenomName, fmt.Sprint(msg.Denom)),
			sdk.NewAttribute(types.AttributeKeyDenomAmount, fmt.Sprint(msg.Amount)),
		),
	)

	return &types.MsgBurnTokensResponse{}, nil
}
