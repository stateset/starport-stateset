package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateAgreement{}

type MsgCreateAgreement struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Number string `json:"number" yaml:"number"`
  Name string `json:"name" yaml:"name"`
  Type string `json:"type" yaml:"type"`
  Status string `json:"status" yaml:"status"`
  StartDate string `json:"startDate" yaml:"startDate"`
  EndDate string `json:"endDate" yaml:"endDate"`
  Total string `json:"total" yaml:"total"`
  Hash string `json:"hash" yaml:"hash"`
}

func NewMsgCreateAgreement(creator sdk.AccAddress, number string, name string, type string, status string, startDate string, endDate string, total string, hash string) MsgCreateAgreement {
  return MsgCreateAgreement{
    ID: uuid.New().String(),
		Creator: creator,
    Number: number,
    Name: name,
    Type: type,
    Status: status,
    StartDate: startDate,
    EndDate: endDate,
    Total: total,
    Hash: hash,
	}
}

func (msg MsgCreateAgreement) Route() string {
  return RouterKey
}

func (msg MsgCreateAgreement) Type() string {
  return "CreateAgreement"
}

func (msg MsgCreateAgreement) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateAgreement) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateAgreement) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}