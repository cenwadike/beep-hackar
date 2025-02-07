package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSendIntentPacket{}

func NewMsgSendIntentPacket(
	creator string,
	port string,
	channelID string,
	amount uint64,
	timeoutTimestamp uint64,
	intentType string,
	memo string,
	targetChain string,
	minOutput uint64,
) *MsgSendIntentPacket {
	return &MsgSendIntentPacket{
		Creator:     creator,
		Port:        port,
		ChannelID:   channelID,
		IntentType:  intentType,
		Memo:        memo,
		TargetChain: targetChain,
		MinOutput:   minOutput,
		Amount:      amount,
	}
}

func (msg *MsgSendIntentPacket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	return nil
}
