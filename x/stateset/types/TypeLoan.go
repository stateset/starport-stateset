package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Loan struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Number string `json:"number" yaml:"number"`
  Name string `json:"name" yaml:"name"`
  Reason string `json:"reason" yaml:"reason"`
  Status string `json:"status" yaml:"status"`
  AmountDue string `json:"amountDue" yaml:"amountDue"`
  AmountPaid string `json:"amountPaid" yaml:"amountPaid"`
  Total string `json:"total" yaml:"total"`
}