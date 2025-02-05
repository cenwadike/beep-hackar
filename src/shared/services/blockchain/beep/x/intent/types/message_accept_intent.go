package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAcceptIntent{}

func NewMsgAcceptIntent(executor string, id uint64) *MsgAcceptIntent {
	return &MsgAcceptIntent{
		Executor: executor,
		Id:       id,
	}
}

func (msg *MsgAcceptIntent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Executor)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid executor address (%s)", err)
	}
	return nil
}
