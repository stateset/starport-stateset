package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Number struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Name string `json:"name" yaml:"name"`
  Reason string `json:"reason" yaml:"reason"`
  AmountDue string `json:"amountDue" yaml:"amountDue"`
  AmountPaid string `json:"amountPaid" yaml:"amountPaid"`
  Total string `json:"total" yaml:"total"`
}