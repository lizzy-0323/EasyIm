package business

import (
	"context"
	"go-im/internal/business/api"
	"go-im/pkg/interceptor"
	"go-im/pkg/logger"
	"go-im/pkg/protocol/pb"
	"go-im/pkg/urlwhitelist"
	"net"
	"os"
	"os/signal"
	"syscall"

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

	// start business server
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.NewInterceptor("business_interceptor", urlwhitelist.Business)))

	// 监听服务关闭信号，服务平滑重启
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		s := <-c
		logger.Logger.Info("server stop", zap.Any("signal", s))
		server.GracefulStop()
	}()

	pb.RegisterBusinessIntServer(server, &api.BusinessIntServer{})
	pb.RegisterBusinessExtServer(server, &api.BusinessExtServer{})
	listen, err := net.Listen("tcp", rpcServerAddress)
	if err != nil {
		panic(err)
	}

	log.Info("business rpc start", zap.String("address", rpcServerAddress))
	err = server.Serve(listen)
	if err != nil {
		log.Error("serve error", zap.Error(err))
	}
	return nil
}
