package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreatePoll{}

// MsgCreatePoll and stuff
type MsgCreatePoll struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Title   string         `json:"title" yaml:"title"`
	Options []string       `json:"options" yaml:"options"`
}

// NewMsgCreatePoll and stuff
func NewMsgCreatePoll(creator sdk.AccAddress, title string, options []string) MsgCreatePoll {
	return MsgCreatePoll{
		Creator: creator,
		Title:   title,
		Options: options,
	}
}

// Route and stuff
func (msg MsgCreatePoll) Route() string {
	return RouterKey
}

// Type and stuff
func (msg MsgCreatePoll) Type() string {
	return "CreatePoll"
}

// GetSigners and stuff
func (msg MsgCreatePoll) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes and stuff
func (msg MsgCreatePoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic and stuff
func (msg MsgCreatePoll) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
