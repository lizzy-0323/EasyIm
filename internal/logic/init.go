package logic

import (
	"context"
	"go-im/internal/logic/api"
	"go-im/pkg/interceptor"
	"go-im/pkg/protocol/pb"
	"go-im/pkg/urlwhitelist"
	"net"

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

func Start(ctx context.Context, rpcServerAddress string) error {
	log = initLogger("debug")

	// start logic server
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.NewInterceptor("logic_interceptor", urlwhitelist.Logic)))

	pb.RegisterLogicIntServer(server, &api.LogicIntServer{})
	pb.RegisterLogicExtServer(server, &api.LogicExtServer{})

	listen, err := net.Listen("tcp", rpcServerAddress)
	if err != nil {
		panic(err)
	}

	log.Info("rpc service start", zap.String("address", rpcServerAddress))
	err = server.Serve(listen)
	if err != nil {
		log.Error("serve error", zap.Error(err))
	}
	return nil
}
