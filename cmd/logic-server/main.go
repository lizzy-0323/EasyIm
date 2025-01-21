package main

import (
	"go-im/config"
	"go-im/internal/logic"
	"go-im/internal/logic/domain/device"
	"go-im/internal/logic/domain/message"
	"go-im/internal/logic/proxy"

	"github.com/spf13/cobra"
)

func init() {
	proxy.MessageProxy = message.App
	proxy.DeviceProxy = device.App
}

// start logic server
func main() {
	logicCmd := NewLogicServerCmd()
	err := logicCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func NewLogicServerCmd() *cobra.Command {
	var rpcServerAddress string
	logicServerCmd := &cobra.Command{
		Use:   "logic-server",
		Short: "Start the logic server",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := logic.Start(cmd.Context(), rpcServerAddress)
			return err
		},
	}
	logicServerCmd.Flags().StringVarP(&rpcServerAddress, "rpc-address", "r", config.Config.LogicRPCListenAddr, "logic rpc server address")
	return logicServerCmd
}
