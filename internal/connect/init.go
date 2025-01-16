package connect

import (
	"context"
	"go-im/pkg/protocol/pb"
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

// start msg gateway server
func Start(ctx context.Context, wsAddress string, rpcServerAddress string, version string) error {
	log = initLogger(version)

	// Start websocket server
	ws := NewWsServer(wsAddress)

	go func() {
		ws.Run()
	}()

	// Start rpc server
	rpcServer := grpc.NewServer()

	pb.RegisterConnectIntServer(rpcServer, &ConnIntServer{})
	listener, err := net.Listen("tcp", rpcServerAddress)
	if err != nil {
		panic(err)
	}

	log.Info("rpc service start", zap.String("address", rpcServerAddress))
	err = rpcServer.Serve(listener)
	if err != nil {
		return err
	}
	return nil
}
