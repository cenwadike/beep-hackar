package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBurnTokens{}

func NewMsgBurnTokens(signer string, amount int32) *MsgBurnTokens {
	return &MsgBurnTokens{
		Signer: signer,
		Amount: amount,
	}
}

func (msg *MsgBurnTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}
