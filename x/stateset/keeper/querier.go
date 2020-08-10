package keeper

import (
  // this line is used by starport scaffolding
	"github.com/stateset/stateset/x/stateset/types"
		
	
		
	
		
	
		
	
		
	
		
	
		
	
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for stateset clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListCase:
			return listCase(ctx, k)
		case types.QueryListLoan:
			return listLoan(ctx, k)
		case types.QueryListContact:
			return listContact(ctx, k)
		case types.QueryListInvoice:
			return listInvoice(ctx, k)
		case types.QueryListPurchaseorder:
			return listPurchaseorder(ctx, k)
		case types.QueryListAgreement:
			return listAgreement(ctx, k)
		case types.QueryListProposal:
			return listProposal(ctx, k)
		case types.QueryListUser:
			return listUser(ctx, k)
		case types.QueryListNumber:
			return listNumber(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown stateset query endpoint")
		}
	}
}