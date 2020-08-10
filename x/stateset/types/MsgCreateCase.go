package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateCase{}

type MsgCreateCase struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Name string `json:"name" yaml:"name"`
  Number string `json:"number" yaml:"number"`
  Description string `json:"description" yaml:"description"`
  Status string `json:"status" yaml:"status"`
}

func NewMsgCreateCase(creator sdk.AccAddress, name string, number string, description string, status string) MsgCreateCase {
  return MsgCreateCase{
    ID: uuid.New().String(),
		Creator: creator,
    Name: name,
    Number: number,
    Description: description,
    Status: status,
	}
}

func (msg MsgCreateCase) Route() string {
  return RouterKey
}

func (msg MsgCreateCase) Type() string {
  return "CreateCase"
}

func (msg MsgCreateCase) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateCase) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateCase) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}