package stateset

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/keeper"
	"github.com/stateset/stateset/x/stateset/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
    // this line is used by starport scaffolding
		case types.MsgCreateContact:
			return handleMsgCreateContact(ctx, k, msg)
		case types.MsgCreateInvoice:
			return handleMsgCreateInvoice(ctx, k, msg)
		case types.MsgCreatePurchaseorder:
			return handleMsgCreatePurchaseorder(ctx, k, msg)
		case types.MsgCreateAgreement:
			return handleMsgCreateAgreement(ctx, k, msg)
		case types.MsgCreateProposal:
			return handleMsgCreateProposal(ctx, k, msg)
		case types.MsgCreateUser:
			return handleMsgCreateUser(ctx, k, msg)
		case types.MsgCreateNumber:
			return handleMsgCreateNumber(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
