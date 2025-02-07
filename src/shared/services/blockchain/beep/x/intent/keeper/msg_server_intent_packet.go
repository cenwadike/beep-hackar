package keeper

import (
	"context"
	"time"

	"beep/x/intent/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
)

func (k msgServer) SendIntentPacket(goCtx context.Context, msg *types.MsgSendIntentPacket) (*types.MsgSendIntentPacketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Construct the packet
	var packet types.IntentPacketPacketData

	packet.IntentType = msg.IntentType
	packet.Memo = msg.Memo
	packet.TargetChain = msg.TargetChain
	packet.MinOutput = msg.MinOutput
	packet.Creator = msg.Creator
	packet.InputToken = msg.InputToken
	packet.OutputToken = msg.OutputToken
	packet.Amount = msg.Amount

	// Construct IBC transfer message with a timeout of one minute
	currentBlockTime := ctx.BlockTime()
	timeoutTimestamp := currentBlockTime.Add(1 * time.Minute).UnixNano()

	currentBlockHeight := ctx.BlockHeight()
	timeoutTimesHeight := currentBlockHeight + 10 // ~1 minute

	// Transmit the packet
	_, err := k.TransmitIntentPacketPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.NewHeight(0, uint64(timeoutTimesHeight)), // Timeout height
		uint64(timeoutTimestamp),
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIntentPacketResponse{}, nil
}
