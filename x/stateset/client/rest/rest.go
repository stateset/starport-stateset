package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers stateset-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/stateset/contact", listContactHandler(cliCtx, "stateset")).Methods("GET")
	r.HandleFunc("/stateset/contact", createContactHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/stateset/invoice", listInvoiceHandler(cliCtx, "stateset")).Methods("GET")
	r.HandleFunc("/stateset/invoice", createInvoiceHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/stateset/purchaseorder", listPurchaseorderHandler(cliCtx, "stateset")).Methods("GET")
	r.HandleFunc("/stateset/purchaseorder", createPurchaseorderHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/stateset/agreement", listAgreementHandler(cliCtx, "stateset")).Methods("GET")
	r.HandleFunc("/stateset/agreement", createAgreementHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/stateset/proposal", listProposalHandler(cliCtx, "stateset")).Methods("GET")
	r.HandleFunc("/stateset/proposal", createProposalHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/stateset/user", listUserHandler(cliCtx, "stateset")).Methods("GET")
	r.HandleFunc("/stateset/user", createUserHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/stateset/number", listNumberHandler(cliCtx, "stateset")).Methods("GET")
	r.HandleFunc("/stateset/number", createNumberHandler(cliCtx)).Methods("POST")
}
