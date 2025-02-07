package keeper

import (
	"errors"
	"fmt"
	"time"

	"beep/x/intent/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
)

// TransmitIntentPacketPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIntentPacketPacket(
	ctx sdk.Context,
	packetData types.IntentPacketPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.ScopedKeeper().GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, errorsmod.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return 0, errorsmod.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %s", err)
	}

	return k.ibcKeeperFn().ChannelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvIntentPacketPacket processes packet reception
func (k Keeper) OnRecvIntentPacketPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IntentPacketPacketData) (packetAck types.IntentPacketPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// Get source channel and port
	_, found := k.ibcKeeperFn().ChannelKeeper.GetChannel(ctx, packet.DestinationPort, packet.DestinationChannel)
	if !found {
		errMsg := fmt.Sprintf("port ID (%s) channel ID (%s)", packet.DestinationPort, packet.DestinationChannel)
		return packetAck, errorsmod.Wrap(channeltypes.ErrChannelNotFound, errMsg)
	}

	senderAddr, err := sdk.AccAddressFromBech32(data.Creator)
	if err != nil {
		return packetAck, err
	}

	// reciever is module
	receiverAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	// get expiration time
	currentBlockTime := ctx.BlockTime()
	expiryHeight := currentBlockTime.Add(1 * time.Minute).Unix()

	// Create a new intent
	intent := types.Intents{
		Id:           k.GetIntentsCount(ctx),
		Creator:      data.Creator,
		ActionType:   data.IntentType,
		InputToken:   data.InputToken,
		OutputToken:  data.OutputToken,
		TargetChain:  data.TargetChain,
		MinOutput:    uint64(data.MinOutput),
		Status:       "OPEN",
		ExpiryHeight: uint64(expiryHeight), // Example expiry height
	}

	// Process the attached IBC token
	if err := k.ProcessIBCToken(ctx, senderAddr, receiverAddr, data.InputToken, int64(data.Amount), data.Memo, packet); err != nil {
		return packetAck, err
	}

	k.AppendIntents(ctx, intent)

	return packetAck, nil
}

// ProcessIBCToken processes the attached IBC token in the packet
func (k Keeper) ProcessIBCToken(ctx sdk.Context, sender, receiver sdk.AccAddress, inputToken string, amount int64, memo string, packet channeltypes.Packet) error {
	// Construct IBC transfer message with a timeout of one minute
	currentBlockTime := ctx.BlockTime()
	timeoutTimestamp := currentBlockTime.Add(1 * time.Minute).UnixNano()

	currentBlockHeight := ctx.BlockHeight()
	timeoutTimesHeight := currentBlockHeight + 10 // ~1 minute

	// Construct IBC transfer message
	transferMsg := ibctransfertypes.NewMsgTransfer(
		"transfer",  // Source port
		"channel-0", // Source channel
		sdk.NewCoin(inputToken, math.NewInt(amount)), // Amount
		sender.String(),
		receiver.String(),
		clienttypes.NewHeight(0, uint64(timeoutTimesHeight)), // Timeout height
		uint64(timeoutTimestamp),                             // Timeout timestamp
		memo,
	)

	if _, err := k.transferKeeper.Transfer(ctx, transferMsg); err != nil {
		return err
	}

	return nil
}

// OnAcknowledgementIntentPacketPacket responds to the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIntentPacketPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IntentPacketPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IntentPacketPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIntentPacketPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIntentPacketPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IntentPacketPacketData) error {

	// TODO: packet timeout logic

	return nil
}
