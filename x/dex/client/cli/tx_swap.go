package cli

import (
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSwap() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap [amount-in] [tokenA] [tokenB] [token-in] [slippage-tolerance] [minOut] [receiver]",
		Short: "Broadcast message swap",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmountIn := args[0]

			argAmountInDec, err := sdk.NewDecFromStr(argAmountIn)

			if err != nil {
				return err
			}

			argTokenA := args[1]
			argTokenB := args[2]
			argTokenIn := args[3]
			argminOut := args[4]

			argminOutDec, err := sdk.NewDecFromStr(argminOut)

			if err != nil {
				return err
			}

			argReceiver := args[5]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSwap(
				clientCtx.GetFromAddress().String(),
				argTokenA,
				argTokenB,
				argAmountInDec,
				argTokenIn,
				argminOutDec,
				argReceiver,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
