package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateLoan(ctx sdk.Context, loan types.Loan) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.LoanPrefix + loan.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(loan)
	store.Set(key, value)
}

func listLoan(ctx sdk.Context, k Keeper) ([]byte, error) {
  var loanList []types.Loan
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.LoanPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var loan types.Loan
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &loan)
    loanList = append(loanList, loan)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, loanList)
  return res, nil
}