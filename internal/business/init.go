package business

import (
	"context"
	"go-im/internal/business/api"
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

func Start(ctx context.Context) error {
	log = initLogger("debug")

	// start business server
	server := grpc.NewServer()

	pb.RegisterBusinessIntServer(server, &api.BusinessIntServer{})
	pb.RegisterBusinessExtServer(server, &api.BusinessExtServer{})
	return nil
}
