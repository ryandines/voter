package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePoll{}

// NewMsgCreatePoll does stuff
func NewMsgCreatePoll(creator string, title string, options []string) *MsgCreatePoll {
	return &MsgCreatePoll{
		Creator: creator,
		Title:   title,
		Options: options,
	}
}

// Route does stuff
func (msg *MsgCreatePoll) Route() string {
	return RouterKey
}

// Type does stuff
func (msg *MsgCreatePoll) Type() string {
	return "CreatePoll"
}

// GetSigners does stuff
func (msg *MsgCreatePoll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes does stuff
func (msg *MsgCreatePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic does stuff
func (msg *MsgCreatePoll) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePoll{}

// NewMsgUpdatePoll does stuff
func NewMsgUpdatePoll(creator string, id string, title string, options []string) *MsgUpdatePoll {
	return &MsgUpdatePoll{
		Id:      id,
		Creator: creator,
		Title:   title,
		Options: options,
	}
}

// Route does stuff
func (msg *MsgUpdatePoll) Route() string {
	return RouterKey
}

// Type does stuff
func (msg *MsgUpdatePoll) Type() string {
	return "UpdatePoll"
}

// GetSigners does stuff
func (msg *MsgUpdatePoll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes does stuff
func (msg *MsgUpdatePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic does stuff
func (msg *MsgUpdatePoll) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreatePoll{}

// NewMsgDeletePoll does stuff
func NewMsgDeletePoll(creator string, id string) *MsgDeletePoll {
	return &MsgDeletePoll{
		Id:      id,
		Creator: creator,
	}
}

// Route does stuff
func (msg *MsgDeletePoll) Route() string {
	return RouterKey
}

// Type does stuff
func (msg *MsgDeletePoll) Type() string {
	return "DeletePoll"
}

// GetSigners does stuff
func (msg *MsgDeletePoll) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes does stuff
func (msg *MsgDeletePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic does stuff
func (msg *MsgDeletePoll) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
