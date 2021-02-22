package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/ryandines/voter/x/voter/types"
)

// CmdCreatePoll does stuff
func CmdCreatePoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-poll [title] [options]",
		Short: "Creates a new poll",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsTitle := string(args[0])
			argsOptions := args[1:]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePoll(clientCtx.GetFromAddress().String(), string(argsTitle), argsOptions)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdUpdatePoll does stuff
func CmdUpdatePoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-poll [id] [title] [options]",
		Short: "Update a poll",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsTitle := string(args[1])
			argsOptions := args[2:]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePoll(clientCtx.GetFromAddress().String(), id, string(argsTitle), argsOptions)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// CmdDeletePoll does stuff
func CmdDeletePoll() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-poll [id] [title] [options]",
		Short: "Delete a poll by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePoll(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
