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

func GetCmdCreateLoan(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-loan [number] [name] [reason] [status] [amountDue] [amountPaid] [total]",
		Short: "Creates a new loan",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsNumber := string(args[0])
      argsName := string(args[1])
      argsReason := string(args[2])
      argsStatus := string(args[3])
      argsAmountDue := string(args[4])
      argsAmountPaid := string(args[5])
      argsTotal := string(args[6])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateLoan(cliCtx.GetFromAddress(), argsNumber, argsName, argsReason, argsStatus, argsAmountDue, argsAmountPaid, argsTotal)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}