package rpc

import (
	"go-im/config"
	"go-im/pkg/protocol/pb"
)

var (
	connectIntClient  pb.ConnectIntClient
	logicIntClient    pb.LogicIntClient
	businessIntClient pb.BusinessIntClient
)

func GetConnectIntClient() pb.ConnectIntClient {
	if connectIntClient == nil {
		connectIntClient = config.Config.ConnectIntClientBuilder()
	}
	return connectIntClient
}

func GetLogicIntClient() pb.LogicIntClient {
	if logicIntClient == nil {
		logicIntClient = config.Config.LogicIntClientBuilder()
	}
	return logicIntClient
}

func GetBusinessIntClient() pb.BusinessIntClient {
	if businessIntClient == nil {
		businessIntClient = config.Config.BusinessIntClientBuilder()
	}
	return businessIntClient
}
