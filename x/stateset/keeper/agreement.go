package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateAgreement(ctx sdk.Context, agreement types.Agreement) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.AgreementPrefix + agreement.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(agreement)
	store.Set(key, value)
}

func listAgreement(ctx sdk.Context, k Keeper) ([]byte, error) {
  var agreementList []types.Agreement
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.AgreementPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var agreement types.Agreement
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &agreement)
    agreementList = append(agreementList, agreement)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, agreementList)
  return res, nil
}