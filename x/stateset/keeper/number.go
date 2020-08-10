package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateNumber(ctx sdk.Context, number types.Number) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.NumberPrefix + number.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(number)
	store.Set(key, value)
}

func listNumber(ctx sdk.Context, k Keeper) ([]byte, error) {
  var numberList []types.Number
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.NumberPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var number types.Number
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &number)
    numberList = append(numberList, number)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, numberList)
  return res, nil
}