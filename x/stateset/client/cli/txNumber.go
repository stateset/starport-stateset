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

func GetCmdCreateNumber(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-number [name] [reason] [amountDue] [amountPaid] [total]",
		Short: "Creates a new number",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsName := string(args[0])
      argsReason := string(args[1])
      argsAmountDue := string(args[2])
      argsAmountPaid := string(args[3])
      argsTotal := string(args[4])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateNumber(cliCtx.GetFromAddress(), argsName, argsReason, argsAmountDue, argsAmountPaid, argsTotal)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}