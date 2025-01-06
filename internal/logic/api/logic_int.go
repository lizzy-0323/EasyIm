package api

import (
	"context"
	"go-im/pkg/protocol/pb"

	"go-im/internal/logic/domain/device"

	"google.golang.org/protobuf/types/known/emptypb"
)

type LogicIntServer struct {
	pb.UnsafeLogicIntServer
}

// ConnSignIn 设备登录
func (*LogicIntServer) ConnSignIn(ctx context.Context, req *pb.ConnSignInReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{},
		device.App.SignIn(ctx, req.UserId, req.DeviceId, req.Token, req.ConnAddr, req.ClientAddr)
}

// GetDevice 获取设备信息
func (*LogicIntServer) GetDevice(ctx context.Context, req *pb.GetDeviceReq) (*pb.GetDeviceResp, error) {
	device, err := device.App.GetDevice(ctx, req.DeviceId)
	return &pb.GetDeviceResp{Device: device}, err
}

// MessageACK 设备收到消息ack
func (*LogicIntServer) MessageACK(ctx context.Context, req *pb.MessageACKReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// Offline 设备离线
func (*LogicIntServer) Offline(ctx context.Context, req *pb.OfflineReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, device.App.Offline(ctx, req.DeviceId, req.ClientAddr)
}

func (s *LogicIntServer) SubscribeRoom(ctx context.Context, req *pb.SubscribeRoomReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// Push 推送
func (*LogicIntServer) Push(ctx context.Context, req *pb.PushReq) (*pb.PushResp, error) {
	return nil, nil
}

// PushRoom 推送房间
func (s *LogicIntServer) PushRoom(ctx context.Context, req *pb.PushRoomReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// PushAll 全服推送
func (s *LogicIntServer) PushAll(ctx context.Context, req *pb.PushAllReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// ServerStop 服务停止
func (s *LogicIntServer) ServerStop(ctx context.Context, in *pb.ServerStopReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

// Sync 设备同步
func (s *LogicIntServer) Sync(ctx context.Context, in *pb.SyncReq) (*pb.SyncResp, error) {
	return nil, nil
}
