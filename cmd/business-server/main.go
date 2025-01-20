package main

import (
	"go-im/config"
	"go-im/internal/business"

	"github.com/spf13/cobra"
)

func main() {
	businessCmd := NewBusinessServer()
	err := businessCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func NewBusinessServer() *cobra.Command {
	var rpcServerAddress string
	businessServerCmd := &cobra.Command{
		Use:   "business-server",
		Short: "Start the business server",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := business.Start(cmd.Context(), rpcServerAddress)
			return err
		},
	}
	businessServerCmd.Flags().StringVarP(&rpcServerAddress, "rpc-address", "r", config.Config.BusinessRPCListenAddr, "business rpc server address")
	return businessServerCmd
}
