package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Contact struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
  FirstName string `json:"firstName" yaml:"firstName"`
  LastName string `json:"lastName" yaml:"lastName"`
  Email string `json:"email" yaml:"email"`
  Phone string `json:"phone" yaml:"phone"`
  Company string `json:"company" yaml:"company"`
}