package cmd

import (
	msggateway "go-im/internal/msg-gateway"

	"github.com/spf13/cobra"
)

func NewMsgGateWayCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "msg-gateway",
		Short: "Start the message server",
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: use viper to load file
			address := "127.0.0.1"
			port := 8080
			if err := msggateway.Start(cmd.Context(), address, port); err != nil {
				return err
			}
			return nil
		},
	}
}
