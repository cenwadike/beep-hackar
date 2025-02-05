package cli

import (
	"strconv"

	"beep/x/intent/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	channelutils "github.com/cosmos/ibc-go/v8/modules/core/04-channel/client/utils"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

// CmdSendIntentPacket() returns the IntentPacket send packet command.
// This command does not use AutoCLI because it gives a better UX to do not.
func CmdSendIntentPacket() *cobra.Command {
	flagPacketTimeoutTimestamp := "packet-timeout-timestamp"

	cmd := &cobra.Command{
		Use:   "send-intent-packet [src-port] [src-channel] [action-type] [memo] [target-chain] [min-output] [status] [executor] [expiry-height]",
		Short: "Send a intentPacket over IBC",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]

			argActionType := args[2]
			argMemo := args[3]
			argTargetChain := args[4]
			argMinOutput, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}
			argStatus := args[6]
			argExecutor := args[7]
			argExpiryHeight, err := cast.ToUint64E(args[8])
			if err != nil {
				return err
			}

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgSendIntentPacket(creator, srcPort, srcChannel, timeoutTimestamp, argActionType, argMemo, argTargetChain, argMinOutput, argStatus, argExecutor, argExpiryHeight)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
