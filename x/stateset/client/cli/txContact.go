package cli

import (
	"bufio"
  
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/stateset/stateset/x/stateset/types"
)

func GetCmdCreateContact(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-contact [firstName] [lastName] [email] [phone] [company]",
		Short: "Creates a new contact",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsFirstName := string(args[0])
      argsLastName := string(args[1])
      argsEmail := string(args[2])
      argsPhone := string(args[3])
      argsCompany := string(args[4])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateContact(cliCtx.GetFromAddress(), argsFirstName, argsLastName, argsEmail, argsPhone, argsCompany)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}