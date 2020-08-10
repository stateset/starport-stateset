package stateset

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
	"github.com/stateset/stateset/x/stateset/keeper"
)

func handleMsgCreateInvoice(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateInvoice) (*sdk.Result, error) {
	var invoice = types.Invoice{
		Creator: msg.Creator,
		ID:      msg.ID,
    Name: msg.Name,
    Reason: msg.Reason,
    AmountDue: msg.AmountDue,
    AmountPaid: msg.AmountPaid,
    Total: msg.Total,
	}
	k.CreateInvoice(ctx, invoice)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
