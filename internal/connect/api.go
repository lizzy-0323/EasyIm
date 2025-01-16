package connect

import (
	"context"
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
		log.Warn("GetConn warn", zap.Int64("deviceId", req.DeviceId))
		return resp, nil
	}

	if conn.DeviceId != req.DeviceId {
		log.Warn("conn.DeviceId != req.DeviceId", zap.Int64("conn.DeviceId", conn.DeviceId), zap.Int64("req.DeviceId", req.DeviceId))
	}

	conn.Send(pb.PackageType_PT_MESSAGE, 0, req.Message, nil)
	return resp, nil
}
