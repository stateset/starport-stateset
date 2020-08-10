package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateInvoice(ctx sdk.Context, invoice types.Invoice) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.InvoicePrefix + invoice.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(invoice)
	store.Set(key, value)
}

func listInvoice(ctx sdk.Context, k Keeper) ([]byte, error) {
  var invoiceList []types.Invoice
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.InvoicePrefix))
  for ; iterator.Valid(); iterator.Next() {
    var invoice types.Invoice
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &invoice)
    invoiceList = append(invoiceList, invoice)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, invoiceList)
  return res, nil
}