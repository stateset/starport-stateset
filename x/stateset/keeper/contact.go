package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateContact(ctx sdk.Context, contact types.Contact) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ContactPrefix + contact.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(contact)
	store.Set(key, value)
}

func listContact(ctx sdk.Context, k Keeper) ([]byte, error) {
  var contactList []types.Contact
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.ContactPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var contact types.Contact
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &contact)
    contactList = append(contactList, contact)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, contactList)
  return res, nil
}