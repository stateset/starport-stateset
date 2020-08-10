package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
  "github.com/stateset/stateset/x/stateset/types"
)

func GetCmdListLoan(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-loan",
		Short: "list all loan",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListLoan, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Loan\n%s\n", err.Error())
				return nil
			}
			var out []types.Loan
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
