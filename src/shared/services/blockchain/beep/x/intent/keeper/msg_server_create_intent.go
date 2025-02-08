package keeper

import (
	"beep/x/intent/types"
	"context"
	"fmt"
	"strings"
	"time"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
)

func (k msgServer) CreateIntent(goCtx context.Context, msg *types.MsgCreateIntent) (*types.MsgCreateIntentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get expiration time
	currentBlockTime := ctx.BlockTime()
	expiryHeight := currentBlockTime.Add(1 * time.Minute).Unix()

	// Create intent object
	var intent = types.Intents{
		Id:           k.GetIntentsCount(ctx),
		Amount:       uint64(msg.Amount),
		Creator:      msg.Creator,
		ActionType:   msg.IntentType,
		InputToken:   msg.InputToken,
		OutputToken:  msg.OutputToken,
		TargetChain:  msg.TargetChain,
		MinOutput:    uint64(msg.MinOutput),
		Status:       "OPEN",
		ExpiryHeight: uint64(expiryHeight),
	}

	// Parse sender address
	senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// receiver is module
	receiverAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	// Add intents to storage
	k.AppendIntents(ctx, intent)

	// Emit event for intent creation
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreateIntent,
			sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
			sdk.NewAttribute(types.AttributeKeyCreator, intent.Creator),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprint(intent.Amount)),
			sdk.NewAttribute(types.AttributeKeyInputToken, intent.InputToken),
			sdk.NewAttribute(types.AttributeKeyOutputToken, intent.OutputToken),
			sdk.NewAttribute(types.AttributeKeyTargetChain, intent.TargetChain),
			sdk.NewAttribute(types.AttributeKeyMinOutput, fmt.Sprint(intent.MinOutput)),
			sdk.NewAttribute(types.AttributeKeyStatus, intent.Status),
			sdk.NewAttribute(types.AttributeKeyExpiryHeight, fmt.Sprint(intent.ExpiryHeight)),
		),
	)

	// Handle input token (escrow if not IBC)
	if isIBCAsset(intent.InputToken) {
		ibcDenom := getIBCDenom(intent.OutputToken)
		if err := k.verifyAndSendIBCToken(ctx, senderAddr, receiverAddr, ibcDenom, int64(intent.Amount), msg.Memo); err != nil {
			return nil, err
		}

		// Emit IBC transfer event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeEscrowIBCTransfer,
				sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
				sdk.NewAttribute(types.AttributeKeySender, senderAddr.String()),
				sdk.NewAttribute(types.AttributeKeyReceiver, receiverAddr.String()),
				sdk.NewAttribute(types.AttributeKeyIBCDenom, ibcDenom),
				sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprint(intent.Amount)),
			),
		)
	} else {
		if err := k.bankKeeper.SendCoins(ctx, senderAddr, receiverAddr, sdk.NewCoins(sdk.NewCoin(intent.InputToken, math.NewInt(int64(intent.Amount))))); err != nil {
			return nil, err
		}

		// Emit token escrow event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeEscrowNativeTokens,
				sdk.NewAttribute(types.AttributeKeyIntentId, fmt.Sprint(intent.Id)),
				sdk.NewAttribute(types.AttributeKeySender, senderAddr.String()),
				sdk.NewAttribute(types.AttributeKeyReceiver, receiverAddr.String()),
				sdk.NewAttribute(types.AttributeKeyDenom, intent.InputToken),
				sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprint(intent.Amount)),
			),
		)
	}

	return &types.MsgCreateIntentResponse{}, nil
}

func isIBCAsset(token string) bool {
	return strings.HasPrefix(token, "ibc/")
}

func getIBCDenom(token string) string {
	// Extract IBC denom from the token
	parts := strings.Split(token, "/")
	if len(parts) > 1 {
		return parts[1]
	}
	return token
}

func (k Keeper) verifyAndSendIBCToken(ctx sdk.Context, sender sdk.AccAddress, receiver sdk.AccAddress, denom string, amount int64, memo string) error {
	// Verify IBC channel exists and is open
	channel, found := k.ibcKeeperFn().ChannelKeeper.GetChannel(ctx, "intent", "channel-1")
	if !found || channel.State != channeltypes.OPEN {
		return errorsmod.Wrap(sdkerrors.ErrUnknownRequest, "IBC channel not found or not open")
	}

	// Construct IBC transfer message with a timeout of one minute
	currentBlockTime := ctx.BlockTime()
	timeoutTimestamp := currentBlockTime.Add(1 * time.Minute).UnixNano()

	currentBlockHeight := ctx.BlockHeight()
	timeoutTimesHeight := currentBlockHeight + 10 // ~1 minute

	// Construct IBC transfer message
	transferMsg := ibctransfertypes.NewMsgTransfer(
		"transfer",                              // Source port
		"channel-0",                             // Source channel
		sdk.NewCoin(denom, math.NewInt(amount)), // Amount
		sender.String(),
		receiver.String(),
		clienttypes.NewHeight(0, uint64(timeoutTimesHeight)), // Timeout height
		uint64(timeoutTimestamp),                             // Timeout timestamp
		memo,
	)

	// Send the transfer message
	_, err := k.transferKeeper.Transfer(ctx, transferMsg)
	if err != nil {
		// Emit IBC transfer failure event
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeIBCTransferFailure,
				sdk.NewAttribute(types.AttributeKeySender, sender.String()),
				sdk.NewAttribute(types.AttributeKeyReceiver, receiver.String()),
				sdk.NewAttribute(types.AttributeKeyDenom, denom),
				sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprint(amount)),
				sdk.NewAttribute(types.AttributeKeyError, err.Error()),
			),
		)
		return err
	}

	return nil
}
