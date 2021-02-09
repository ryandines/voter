package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetPoll{}

// MsgSetPoll and stuff
type MsgSetPoll struct {
	ID      string         `json:"id" yaml:"id"`
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Title   string         `json:"title" yaml:"title"`
	Options []string       `json:"options" yaml:"options"`
}

// NewMsgSetPoll and stuff
func NewMsgSetPoll(creator sdk.AccAddress, id string, title string, options []string) MsgSetPoll {
	return MsgSetPoll{
		ID:      id,
		Creator: creator,
		Title:   title,
		Options: options,
	}
}

// Route and stuff
func (msg MsgSetPoll) Route() string {
	return RouterKey
}

// Type and stuff
func (msg MsgSetPoll) Type() string {
	return "SetPoll"
}

// GetSigners and stuff
func (msg MsgSetPoll) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes and stuff
func (msg MsgSetPoll) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic and stuff
func (msg MsgSetPoll) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
