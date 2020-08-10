package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreatePurchaseorder{}

type MsgCreatePurchaseorder struct {
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

func NewMsgCreatePurchaseorder(creator sdk.AccAddress, number string, name string, type string, status string, startDate string, endDate string, total string, hash string) MsgCreatePurchaseorder {
  return MsgCreatePurchaseorder{
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

func (msg MsgCreatePurchaseorder) Route() string {
  return RouterKey
}

func (msg MsgCreatePurchaseorder) Type() string {
  return "CreatePurchaseorder"
}

func (msg MsgCreatePurchaseorder) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreatePurchaseorder) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreatePurchaseorder) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}