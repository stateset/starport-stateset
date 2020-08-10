package stateset

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
	"github.com/stateset/stateset/x/stateset/keeper"
)

func handleMsgCreateContact(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateContact) (*sdk.Result, error) {
	var contact = types.Contact{
		Creator: msg.Creator,
		ID:      msg.ID,
    FirstName: msg.FirstName,
    LastName: msg.LastName,
    Email: msg.Email,
    Phone: msg.Phone,
    Company: msg.Company,
	}
	k.CreateContact(ctx, contact)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
