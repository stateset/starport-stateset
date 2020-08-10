package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateContact{}

type MsgCreateContact struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  FirstName string `json:"firstName" yaml:"firstName"`
  LastName string `json:"lastName" yaml:"lastName"`
  Email string `json:"email" yaml:"email"`
  Phone string `json:"phone" yaml:"phone"`
  Company string `json:"company" yaml:"company"`
}

func NewMsgCreateContact(creator sdk.AccAddress, firstName string, lastName string, email string, phone string, company string) MsgCreateContact {
  return MsgCreateContact{
    ID: uuid.New().String(),
		Creator: creator,
    FirstName: firstName,
    LastName: lastName,
    Email: email,
    Phone: phone,
    Company: company,
	}
}

func (msg MsgCreateContact) Route() string {
  return RouterKey
}

func (msg MsgCreateContact) Type() string {
  return "CreateContact"
}

func (msg MsgCreateContact) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateContact) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateContact) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}