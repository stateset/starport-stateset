package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateProposal(ctx sdk.Context, proposal types.Proposal) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ProposalPrefix + proposal.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(proposal)
	store.Set(key, value)
}

func listProposal(ctx sdk.Context, k Keeper) ([]byte, error) {
  var proposalList []types.Proposal
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.ProposalPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var proposal types.Proposal
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &proposal)
    proposalList = append(proposalList, proposal)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, proposalList)
  return res, nil
}