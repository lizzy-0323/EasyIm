package main

import "go-im/pkg/common/cmd"

// start msg gateway server
func main() {
	if err := cmd.NewMsgGateWayCmd().Execute(); err != nil {
		panic(err)
	}
}
