package stateset

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
	"github.com/stateset/stateset/x/stateset/keeper"
)

func handleMsgCreateCase(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateCase) (*sdk.Result, error) {
	var case = types.Case{
		Creator: msg.Creator,
		ID:      msg.ID,
    Name: msg.Name,
    Number: msg.Number,
    Description: msg.Description,
    Status: msg.Status,
	}
	k.CreateCase(ctx, case)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
