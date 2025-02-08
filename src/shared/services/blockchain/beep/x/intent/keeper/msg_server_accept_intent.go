package keeper

import (
	"beep/x/intent/types"
	"context"
	"fmt"

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

	// Emit event for intent acceptance start
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAcceptIntent,
			sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
			sdk.NewAttribute(types.AttributeKeyCreator, intent.Creator),
			sdk.NewAttribute(types.AttributeKeyExecutor, msg.Executor),
			sdk.NewAttribute(types.AttributeKeyInputToken, intent.InputToken),
			sdk.NewAttribute(types.AttributeKeyOutputToken, intent.OutputToken),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprint(intent.Amount)),
			sdk.NewAttribute(types.AttributeKeyMinOutput, fmt.Sprint(intent.MinOutput)),
		),
	)

	if isIBCAsset(intent.OutputToken) {
		ibcDenom := getIBCDenom(intent.OutputToken)

		// Transfer output token to escrow
		if err := k.verifyAndSendIBCToken(ctx, executor, moduleAddr, ibcDenom, int64(intent.MinOutput), ""); err != nil {
			// Emit failure event
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeAcceptIntentFailure,
					sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
					sdk.NewAttribute(types.AttributeKeyError, err.Error()),
					sdk.NewAttribute(types.AttributeKeyStage, "executor_to_escrow"),
				),
			)
			return nil, err
		}

		// Emit successful escrow event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeExecutorEscrow,
				sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
				sdk.NewAttribute(types.AttributeKeySender, executor.String()),
				sdk.NewAttribute(types.AttributeKeyReceiver, moduleAddr.String()),
				sdk.NewAttribute(types.AttributeKeyIBCDenom, ibcDenom),
				sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprint(intent.MinOutput)),
			),
		)

		// Transfer output token to user
		if err := k.verifyAndSendIBCToken(ctx, moduleAddr, creator, ibcDenom, int64(intent.MinOutput), ""); err != nil {
			// Emit failure event
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeAcceptIntentFailure,
					sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
					sdk.NewAttribute(types.AttributeKeyError, err.Error()),
					sdk.NewAttribute(types.AttributeKeyStage, "escrow_to_creator"),
				),
			)
			return nil, err
		}

		// Emit successful transfer to creator event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeCreatorPayout,
				sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
				sdk.NewAttribute(types.AttributeKeySender, moduleAddr.String()),
				sdk.NewAttribute(types.AttributeKeyReceiver, creator.String()),
				sdk.NewAttribute(types.AttributeKeyIBCDenom, ibcDenom),
				sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprint(intent.MinOutput)),
			),
		)
	} else {
		// Transfer output token to escrow
		if err := k.bankKeeper.SendCoins(ctx, executor, moduleAddr, sdk.NewCoins(sdk.NewCoin(intent.OutputToken, math.NewInt(int64(intent.Amount))))); err != nil {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeAcceptIntentFailure,
					sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
					sdk.NewAttribute(types.AttributeKeyError, err.Error()),
					sdk.NewAttribute(types.AttributeKeyStage, "executor_to_escrow"),
				),
			)
			return nil, err
		}

		// Emit successful escrow event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeExecutorEscrow,
				sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
				sdk.NewAttribute(types.AttributeKeySender, executor.String()),
				sdk.NewAttribute(types.AttributeKeyReceiver, moduleAddr.String()),
				sdk.NewAttribute(types.AttributeKeyDenom, intent.OutputToken),
				sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprint(intent.Amount)),
			),
		)

		// Transfer output token to user
		if err := k.bankKeeper.SendCoins(ctx, moduleAddr, creator, sdk.NewCoins(sdk.NewCoin(intent.OutputToken, math.NewInt(int64(intent.Amount))))); err != nil {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeAcceptIntentFailure,
					sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
					sdk.NewAttribute(types.AttributeKeyError, err.Error()),
					sdk.NewAttribute(types.AttributeKeyStage, "escrow_to_creator"),
				),
			)
			return nil, err
		}

		// Emit successful transfer to creator event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeCreatorPayout,
				sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
				sdk.NewAttribute(types.AttributeKeySender, moduleAddr.String()),
				sdk.NewAttribute(types.AttributeKeyReceiver, creator.String()),
				sdk.NewAttribute(types.AttributeKeyDenom, intent.OutputToken),
				sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprint(intent.Amount)),
			),
		)
	}

	// Transfer input token to executor
	if isIBCAsset(intent.InputToken) {
		ibcDenom := getIBCDenom(intent.InputToken)
		if err := k.verifyAndSendIBCToken(ctx, moduleAddr, executor, ibcDenom, int64(intent.Amount), ""); err != nil {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeAcceptIntentFailure,
					sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
					sdk.NewAttribute(types.AttributeKeyError, err.Error()),
					sdk.NewAttribute(types.AttributeKeyStage, "escrow_to_executor"),
				),
			)
			return nil, err
		}

		// Emit successful executor payout event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeExecutorPayout,
				sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
				sdk.NewAttribute(types.AttributeKeySender, moduleAddr.String()),
				sdk.NewAttribute(types.AttributeKeyReceiver, executor.String()),
				sdk.NewAttribute(types.AttributeKeyIBCDenom, ibcDenom),
				sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprint(intent.Amount)),
			),
		)
	} else {
		if err := k.bankKeeper.SendCoins(ctx, moduleAddr, executor, sdk.NewCoins(sdk.NewCoin(intent.InputToken, math.NewInt(int64(intent.Amount))))); err != nil {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeAcceptIntentFailure,
					sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
					sdk.NewAttribute(types.AttributeKeyError, err.Error()),
					sdk.NewAttribute(types.AttributeKeyStage, "escrow_to_executor"),
				),
			)
			return nil, err
		}

		// Emit successful executor payout event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeExecutorPayout,
				sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
				sdk.NewAttribute(types.AttributeKeySender, moduleAddr.String()),
				sdk.NewAttribute(types.AttributeKeyReceiver, executor.String()),
				sdk.NewAttribute(types.AttributeKeyDenom, intent.InputToken),
				sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprint(intent.Amount)),
			),
		)
	}

	// Update state
	intent.Status = "EXECUTED"
	k.SetIntents(ctx, intent)

	// Emit final success event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeIntentExecuted,
			sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
			sdk.NewAttribute(types.AttributeKeyCreator, intent.Creator),
			sdk.NewAttribute(types.AttributeKeyExecutor, msg.Executor),
			sdk.NewAttribute(types.AttributeKeyStatus, intent.Status),
		),
	)

	return &types.MsgAcceptIntentResponse{}, nil
}
