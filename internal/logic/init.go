package logic

import (
	"go-im/internal/logic/api"
	"go-im/pkg/protocol/pb"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var log *zap.Logger

func initLogger(version string) (logger *zap.Logger) {
	if version == "debug" {
		logger = zap.NewExample()
	} else {
		logger, _ = zap.NewProduction()
	}
	return logger
}

func Start() error {
	log = initLogger("debug")

	// start logic server
	server := grpc.NewServer()

	pb.RegisterLogicIntServer(server, &api.LogicIntServer{})
	return nil
}

func init() {

}
