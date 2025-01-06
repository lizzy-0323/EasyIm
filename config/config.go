package config

import (
	"context"
	"fmt"
	"go-im/pkg/protocol/pb"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

var Config Configuration

var builders = map[string]Builder{}

type Builder interface {
	Build() Configuration
}

type Configuration struct {
	MySQL                string
	RedisHost            string
	RedisPassword        string
	PushRoomSubscribeNum int
	PushAllSubscribeNum  int

	ConnectLocalAddr     string
	ConnectRPCListenAddr string
	ConnectTCPListenAddr string
	ConnectWSListenAddr  string

	LogicRPCListenAddr    string
	BusinessRPCListenAddr string
	FileHTTPListenAddr    string

	ConnectIntClientBuilder  func() pb.ConnectIntClient
	LogicIntClientBuilder    func() pb.LogicIntClient
	BusinessIntClientBuilder func() pb.BusinessIntClient

	Version string
}

type defaultBuilder struct{}

func (d *defaultBuilder) Build() Configuration {
	return Configuration{
		ConnectLocalAddr:     "127.0.0.1:8000",
		ConnectRPCListenAddr: ":8000",
		ConnectTCPListenAddr: ":8001",
		ConnectWSListenAddr:  ":8002",

		LogicRPCListenAddr:    ":8010",
		BusinessRPCListenAddr: ":8020",

		Version: "debug",

		ConnectIntClientBuilder: func() pb.ConnectIntClient {
			conn, err := grpc.NewClient(Config.ConnectRPCListenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, "round_robin")))
			if err != nil {
				panic(err)
			}
			return pb.NewConnectIntClient(conn)
		},
		LogicIntClientBuilder: func() pb.LogicIntClient {
			conn, err := grpc.NewClient(Config.LogicRPCListenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewLogicIntClient(conn)
		},
		BusinessIntClientBuilder: func() pb.BusinessIntClient {
			conn, err := grpc.NewClient(Config.BusinessRPCListenAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewBusinessIntClient(conn)
		},
	}
}

func init() {
	env := os.Getenv("IM_ENV")
	builder, ok := builders[env]
	if !ok {
		builder = new(defaultBuilder)
	}
	Config = builder.Build()
}

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}
