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

func GetCmdCreateCase(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-case [name] [number] [description] [status]",
		Short: "Creates a new case",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsName := string(args[0])
      argsNumber := string(args[1])
      argsDescription := string(args[2])
      argsStatus := string(args[3])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateCase(cliCtx.GetFromAddress(), argsName, argsNumber, argsDescription, argsStatus)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}