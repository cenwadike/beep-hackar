package keeper

import (
	"context"
	"strings"
	"time"

	"beep/x/intent/types"

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

	// Validate input token balance
	var intent = types.Intents{
		Id:           k.GetIntentsCount(ctx),
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

	// reciever is module
	receiverAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	// Add intents to storage
	k.AppendIntents(ctx, intent)

	// Escrow tokens from the intent creator
	// Handle input token (escrow if not IBC)
	if isIBCAsset(intent.InputToken) {
		ibcDenom := getIBCDenom(intent.OutputToken)
		if err := k.verifyAndSendIBCToken(ctx, senderAddr, receiverAddr, ibcDenom, int64(intent.Amount), msg.Memo); err != nil {
			return nil, err
		}
	} else {
		if err := k.bankKeeper.SendCoins(ctx, senderAddr, receiverAddr, sdk.NewCoins(sdk.NewCoin(intent.InputToken, math.NewInt(int64(intent.Amount))))); err != nil {
			return nil, err
		}
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
		return err
	}

	return nil
}
