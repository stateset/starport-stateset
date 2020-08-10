package stateset

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
	"github.com/stateset/stateset/x/stateset/keeper"
)

func handleMsgCreatePurchaseorder(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreatePurchaseorder) (*sdk.Result, error) {
	var purchaseorder = types.Purchaseorder{
		Creator: msg.Creator,
		ID:      msg.ID,
    Number: msg.Number,
    Name: msg.Name,
    Type: msg.Type,
    Status: msg.Status,
    StartDate: msg.StartDate,
    EndDate: msg.EndDate,
    Total: msg.Total,
    Hash: msg.Hash,
	}
	k.CreatePurchaseorder(ctx, purchaseorder)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
