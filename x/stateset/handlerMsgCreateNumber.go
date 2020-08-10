package stateset

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
	"github.com/stateset/stateset/x/stateset/keeper"
)

func handleMsgCreateNumber(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateNumber) (*sdk.Result, error) {
	var number = types.Number{
		Creator: msg.Creator,
		ID:      msg.ID,
    Name: msg.Name,
    Reason: msg.Reason,
    AmountDue: msg.AmountDue,
    AmountPaid: msg.AmountPaid,
    Total: msg.Total,
	}
	k.CreateNumber(ctx, number)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
