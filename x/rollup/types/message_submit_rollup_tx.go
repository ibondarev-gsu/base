package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitRollupTx{}

func NewMsgSubmitRollupTx(creator string, data string) *MsgSubmitRollupTx {
	return &MsgSubmitRollupTx{
		Creator: creator,
		Data:    data,
	}
}

func (msg *MsgSubmitRollupTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
