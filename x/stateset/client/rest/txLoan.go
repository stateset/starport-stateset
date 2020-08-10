package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stateset/stateset/x/stateset/types"
)

type createLoanRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Number string `json:"number"`
	Name string `json:"name"`
	Reason string `json:"reason"`
	Status string `json:"status"`
	AmountDue string `json:"amountDue"`
	AmountPaid string `json:"amountPaid"`
	Total string `json:"total"`
	
}

func createLoanHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createLoanRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgCreateLoan(creator,  req.Number,  req.Name,  req.Reason,  req.Status,  req.AmountDue,  req.AmountPaid,  req.Total, )
		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
