package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers stateset-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding
	r.HandleFunc("/stateset/user", listUserHandler(cliCtx, "stateset")).Methods("GET")
	r.HandleFunc("/stateset/user", createUserHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/stateset/number", listNumberHandler(cliCtx, "stateset")).Methods("GET")
	r.HandleFunc("/stateset/number", createNumberHandler(cliCtx)).Methods("POST")
}
