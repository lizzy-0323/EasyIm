package main

import (
	"go-im/config"
	msggateway "go-im/internal/msg-gateway"

	"github.com/spf13/cobra"
)

// start msg gateway server
func main() {
	msgCmd := NewMsgGateWayCmd()
	err := msgCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func NewMsgGateWayCmd() *cobra.Command {
	var wsAddress string
	var rpcServerAddress string
	var version string
	msgGatewayCmd := &cobra.Command{
		Use:   "msg-gateway",
		Short: "Start the message gateway",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := msggateway.Start(cmd.Context(), wsAddress, rpcServerAddress, version)
			return err
		},
	}
	msgGatewayCmd.Flags().StringVarP(&wsAddress, "ws-address", "a", config.Config.ConnectWSListenAddr, "msg-gateway listen address")
	msgGatewayCmd.Flags().StringVarP(&rpcServerAddress, "rpc-address", "r", config.Config.ConnectRPCListenAddr, "msg-gateway rpc server address")
	msgGatewayCmd.Flags().StringVarP(&version, "version", "v", config.Config.Version, "version, production or debug")
	return msgGatewayCmd
}
