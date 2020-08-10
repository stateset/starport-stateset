package stateset

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
	"github.com/stateset/stateset/x/stateset/keeper"
)

func handleMsgCreateProposal(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateProposal) (*sdk.Result, error) {
	var proposal = types.Proposal{
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
	k.CreateProposal(ctx, proposal)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
