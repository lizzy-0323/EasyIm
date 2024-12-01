package main

import "go-im/pkg/common/cmd"

// start api gateway
func main() {
	if err := cmd.NewApiGateWay().Execute(); err != nil {
		panic(err)
	}
}
