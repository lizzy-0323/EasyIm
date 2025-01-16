package main

import (
	"go-im/config"
	"go-im/internal/connect"

	"github.com/spf13/cobra"
)

func main() {
	connCmd := NewConnectServerCmd()
	err := connCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func NewConnectServerCmd() *cobra.Command {
	var wsAddress string
	var rpcServerAddress string
	var version string
	connServerCmd := &cobra.Command{
		Use:   "connect-sever",
		Short: "Start the connect server",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := connect.Start(cmd.Context(), wsAddress, rpcServerAddress, version)
			return err
		},
	}
	connServerCmd.Flags().StringVarP(&wsAddress, "ws-address", "a", config.Config.ConnectWSListenAddr, "websocket listen address")
	connServerCmd.Flags().StringVarP(&rpcServerAddress, "rpc-address", "r", config.Config.ConnectRPCListenAddr, "connect rpc server address")
	connServerCmd.Flags().StringVarP(&version, "version", "v", config.Config.Version, "version, production or debug")
	return connServerCmd
}
