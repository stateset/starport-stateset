package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateNumber{}

type MsgCreateNumber struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Name string `json:"name" yaml:"name"`
  Reason string `json:"reason" yaml:"reason"`
  AmountDue string `json:"amountDue" yaml:"amountDue"`
  AmountPaid string `json:"amountPaid" yaml:"amountPaid"`
  Total string `json:"total" yaml:"total"`
}

func NewMsgCreateNumber(creator sdk.AccAddress, name string, reason string, amountDue string, amountPaid string, total string) MsgCreateNumber {
  return MsgCreateNumber{
    ID: uuid.New().String(),
		Creator: creator,
    Name: name,
    Reason: reason,
    AmountDue: amountDue,
    AmountPaid: amountPaid,
    Total: total,
	}
}

func (msg MsgCreateNumber) Route() string {
  return RouterKey
}

func (msg MsgCreateNumber) Type() string {
  return "CreateNumber"
}

func (msg MsgCreateNumber) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateNumber) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateNumber) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}