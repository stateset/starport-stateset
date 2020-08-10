package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateInvoice{}

type MsgCreateInvoice struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Name string `json:"name" yaml:"name"`
  Reason string `json:"reason" yaml:"reason"`
  AmountDue string `json:"amountDue" yaml:"amountDue"`
  AmountPaid string `json:"amountPaid" yaml:"amountPaid"`
  Total string `json:"total" yaml:"total"`
}

func NewMsgCreateInvoice(creator sdk.AccAddress, name string, reason string, amountDue string, amountPaid string, total string) MsgCreateInvoice {
  return MsgCreateInvoice{
    ID: uuid.New().String(),
		Creator: creator,
    Name: name,
    Reason: reason,
    AmountDue: amountDue,
    AmountPaid: amountPaid,
    Total: total,
	}
}

func (msg MsgCreateInvoice) Route() string {
  return RouterKey
}

func (msg MsgCreateInvoice) Type() string {
  return "CreateInvoice"
}

func (msg MsgCreateInvoice) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateInvoice) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateInvoice) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}