package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreatePurchaseorder(ctx sdk.Context, purchaseorder types.Purchaseorder) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PurchaseorderPrefix + purchaseorder.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(purchaseorder)
	store.Set(key, value)
}

func listPurchaseorder(ctx sdk.Context, k Keeper) ([]byte, error) {
  var purchaseorderList []types.Purchaseorder
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.PurchaseorderPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var purchaseorder types.Purchaseorder
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &purchaseorder)
    purchaseorderList = append(purchaseorderList, purchaseorder)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, purchaseorderList)
  return res, nil
}