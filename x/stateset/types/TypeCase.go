package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Case struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  Name string `json:"name" yaml:"name"`
  Number string `json:"number" yaml:"number"`
  Description string `json:"description" yaml:"description"`
  Status string `json:"status" yaml:"status"`
}