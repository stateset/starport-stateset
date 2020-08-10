package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateCase(ctx sdk.Context, case types.Case) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.CasePrefix + case.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(case)
	store.Set(key, value)
}

func listCase(ctx sdk.Context, k Keeper) ([]byte, error) {
  var caseList []types.Case
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.CasePrefix))
  for ; iterator.Valid(); iterator.Next() {
    var case types.Case
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &case)
    caseList = append(caseList, case)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, caseList)
  return res, nil
}