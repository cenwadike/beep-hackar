package keeper

import (
	"context"

	"beep/x/intent/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
)

func (k msgServer) SendIntentPacket(goCtx context.Context, msg *types.MsgSendIntentPacket) (*types.MsgSendIntentPacketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IntentPacketPacketData

	packet.ActionType = msg.ActionType
	packet.Memo = msg.Memo
	packet.TargetChain = msg.TargetChain
	packet.MinOutput = msg.MinOutput
	packet.Status = msg.Status
	packet.Executor = msg.Executor
	packet.ExpiryHeight = msg.ExpiryHeight

	// Transmit the packet
	_, err := k.TransmitIntentPacketPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIntentPacketResponse{}, nil
}
