package connect

import (
	"context"
	"go-im/config"
	"go-im/pkg/interceptor"
	"go-im/pkg/logger"
	"go-im/pkg/protocol/pb"
	"go-im/pkg/rpc"
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

// start msg gateway server
func Start(ctx context.Context, wsAddress string, serverAddress string, version string) error {
	log = initLogger(version)

	// Start websocket server
	ws := NewWsServer(wsAddress)

	go func() {
		ws.Run()
	}()

	// Start Subscribe
	StartSubscribe()

	// Start rpc server
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.NewInterceptor("connect_interceptor", nil)))

	// 监听服务关闭信号，服务平滑重启
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		s := <-c
		logger.Logger.Info("server stop start", zap.Any("signal", s))
		_, _ = rpc.GetLogicIntClient().ServerStop(context.TODO(), &pb.ServerStopReq{ConnAddr: config.Config.ConnectLocalAddr})
		logger.Logger.Info("server stop end")

		server.GracefulStop()
	}()

	pb.RegisterConnectIntServer(server, &ConnIntServer{})
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		panic(err)
	}

	log.Info("rpc service start", zap.String("address", serverAddress))
	err = server.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
