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

func GetCmdCreateAgreement(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-agreement [number] [name] [type] [status] [startDate] [endDate] [total] [hash]",
		Short: "Creates a new agreement",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsNumber := string(args[0])
      argsName := string(args[1])
      argsType := string(args[2])
      argsStatus := string(args[3])
      argsStartDate := string(args[4])
      argsEndDate := string(args[5])
      argsTotal := string(args[6])
      argsHash := string(args[7])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateAgreement(cliCtx.GetFromAddress(), argsNumber, argsName, argsType, argsStatus, argsStartDate, argsEndDate, argsTotal, argsHash)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}