package keeper

import (
	"context"

	"beep/x/intent/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AcceptIntent(goCtx context.Context, msg *types.MsgAcceptIntent) (*types.MsgAcceptIntentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get intent
	intent, found := k.GetIntents(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "id: %d", msg.Id)
	}

	// Ensure intent is open
	if intent.Status != "OPEN" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "Intent already executed")
	}

	// Parse user address
	creator, err := sdk.AccAddressFromBech32(intent.Creator)
	if err != nil {
		return nil, err
	}

	// Parse executor address
	executor, err := sdk.AccAddressFromBech32(msg.Executor)
	if err != nil {
		return nil, err
	}

	// Get escrow address
	moduleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	if isIBCAsset(intent.OutputToken) {
		ibcDenom := getIBCDenom(intent.OutputToken)

		// Transfer output token to escrow
		if err := k.verifyAndSendIBCToken(ctx, executor, moduleAddr, ibcDenom, int64(intent.MinOutput), ""); err != nil {
			return nil, err
		}

		// Transfer output token to user
		if err := k.verifyAndSendIBCToken(ctx, moduleAddr, creator, ibcDenom, int64(intent.MinOutput), ""); err != nil {
			return nil, err
		}
	} else {
		// Transfer output token to escrow
		if err := k.bankKeeper.SendCoins(ctx, executor, moduleAddr, sdk.NewCoins(sdk.NewCoin(intent.OutputToken, math.NewInt(int64(intent.Amount))))); err != nil {
			return nil, err
		}

		// Transfer output token to user
		if err := k.bankKeeper.SendCoins(ctx, moduleAddr, creator, sdk.NewCoins(sdk.NewCoin(intent.OutputToken, math.NewInt(int64(intent.Amount))))); err != nil {
			return nil, err
		}
	}

	// Transfer input token to executor
	if isIBCAsset(intent.InputToken) {
		ibcDenom := getIBCDenom(intent.InputToken)
		if err := k.verifyAndSendIBCToken(ctx, moduleAddr, executor, ibcDenom, int64(intent.Amount), ""); err != nil {
			return nil, err
		}
	} else {
		if err := k.bankKeeper.SendCoins(ctx, moduleAddr, executor, sdk.NewCoins(sdk.NewCoin(intent.InputToken, math.NewInt(int64(intent.Amount))))); err != nil {
			return nil, err
		}
	}

	// Update state
	intent.Status = "EXECUTED"
	k.SetIntents(ctx, intent)

	return &types.MsgAcceptIntentResponse{}, nil
}
