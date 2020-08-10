package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateLoan{}

type MsgCreateLoan struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Number string `json:"number" yaml:"number"`
  Name string `json:"name" yaml:"name"`
  Reason string `json:"reason" yaml:"reason"`
  Status string `json:"status" yaml:"status"`
  AmountDue string `json:"amountDue" yaml:"amountDue"`
  AmountPaid string `json:"amountPaid" yaml:"amountPaid"`
  Total string `json:"total" yaml:"total"`
}

func NewMsgCreateLoan(creator sdk.AccAddress, number string, name string, reason string, status string, amountDue string, amountPaid string, total string) MsgCreateLoan {
  return MsgCreateLoan{
    ID: uuid.New().String(),
		Creator: creator,
    Number: number,
    Name: name,
    Reason: reason,
    Status: status,
    AmountDue: amountDue,
    AmountPaid: amountPaid,
    Total: total,
	}
}

func (msg MsgCreateLoan) Route() string {
  return RouterKey
}

func (msg MsgCreateLoan) Type() string {
  return "CreateLoan"
}

func (msg MsgCreateLoan) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateLoan) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateLoan) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}