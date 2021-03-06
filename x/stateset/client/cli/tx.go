package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stateset/stateset/x/stateset/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	statesetTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	statesetTxCmd.AddCommand(flags.PostCommands(
    // this line is used by starport scaffolding
		GetCmdCreateCase(cdc),
		GetCmdCreateLoan(cdc),
		GetCmdCreateContact(cdc),
		GetCmdCreateInvoice(cdc),
		GetCmdCreatePurchaseorder(cdc),
		GetCmdCreateAgreement(cdc),
		GetCmdCreateProposal(cdc),
		GetCmdCreateUser(cdc),
		GetCmdCreateNumber(cdc),
	)...)

	return statesetTxCmd
}
