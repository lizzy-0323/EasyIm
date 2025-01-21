package connect

import (
	"context"
	"go-im/pkg/grpclib"
	"go-im/pkg/logger"
	"go-im/pkg/protocol/pb"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ConnIntServer struct {
	pb.UnsafeConnectIntServer
}

func (s *ConnIntServer) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*emptypb.Empty, error) {
	resp := &emptypb.Empty{}

	// Get client
	conn := GetConn(req.DeviceId)
	if conn == nil {
		logger.Logger.Warn("GetConn warn", zap.Int64("deviceId", req.DeviceId))
		return resp, nil
	}

	if conn.DeviceId != req.DeviceId {
		logger.Logger.Warn("conn.DeviceId != req.DeviceId", zap.Int64("conn.DeviceId", conn.DeviceId), zap.Int64("req.DeviceId", req.DeviceId))
	}

	conn.Send(pb.PackageType_PT_MESSAGE, grpclib.GetCtxRequestId(ctx), req.Message, nil)
	return resp, nil
}
