package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Purchaseorder struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Number string `json:"number" yaml:"number"`
  Name string `json:"name" yaml:"name"`
  Type string `json:"type" yaml:"type"`
  Status string `json:"status" yaml:"status"`
  StartDate string `json:"startDate" yaml:"startDate"`
  EndDate string `json:"endDate" yaml:"endDate"`
  Total string `json:"total" yaml:"total"`
  Hash string `json:"hash" yaml:"hash"`
}