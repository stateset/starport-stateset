package stateset

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
	"github.com/stateset/stateset/x/stateset/keeper"
)

func handleMsgCreateLoan(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateLoan) (*sdk.Result, error) {
	var loan = types.Loan{
		Creator: msg.Creator,
		ID:      msg.ID,
    Number: msg.Number,
    Name: msg.Name,
    Reason: msg.Reason,
    Status: msg.Status,
    AmountDue: msg.AmountDue,
    AmountPaid: msg.AmountPaid,
    Total: msg.Total,
	}
	k.CreateLoan(ctx, loan)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
