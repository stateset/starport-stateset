package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stateset/stateset/x/stateset/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group stateset queries under a subcommand
	statesetQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	statesetQueryCmd.AddCommand(
		flags.GetCommands(
      // this line is used by starport scaffolding
			GetCmdListContact(queryRoute, cdc),
			GetCmdListInvoice(queryRoute, cdc),
			GetCmdListPurchaseorder(queryRoute, cdc),
			GetCmdListAgreement(queryRoute, cdc),
			GetCmdListProposal(queryRoute, cdc),
			GetCmdListUser(queryRoute, cdc),
			GetCmdListNumber(queryRoute, cdc),
		)...,
	)

	return statesetQueryCmd
}
