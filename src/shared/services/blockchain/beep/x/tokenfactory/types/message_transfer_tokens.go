package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgTransferTokens{}

func NewMsgTransferTokens(denom string, owner string, sender string, recipient string, amount int32) *MsgTransferTokens {
	return &MsgTransferTokens{
		Denom:     denom,
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
}

func (msg *MsgTransferTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
