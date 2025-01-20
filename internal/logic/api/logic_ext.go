package api

import (
	"context"
	"go-im/internal/logic/domain/device"
	"go-im/pkg/protocol/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type LogicExtServer struct {
	pb.UnsafeLogicExtServer
}

// RegisterDevice 注册设备
func (*LogicExtServer) RegisterDevice(ctx context.Context, in *pb.RegisterDeviceReq) (*pb.RegisterDeviceResp, error) {
	deviceId, err := device.App.Register(ctx, in)
	return &pb.RegisterDeviceResp{DeviceId: deviceId}, err
}

func (s *LogicExtServer) PushRoom(ctx context.Context, req *pb.PushRoomReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (*LogicExtServer) SendMessageToFriend(ctx context.Context, in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	return nil, nil
}

func (s *LogicExtServer) AddFriend(ctx context.Context, in *pb.AddFriendReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *LogicExtServer) AgreeAddFriend(ctx context.Context, in *pb.AgreeAddFriendReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *LogicExtServer) SetFriend(ctx context.Context, req *pb.SetFriendReq) (*pb.SetFriendResp, error) {
	return &pb.SetFriendResp{}, nil
}

func (s *LogicExtServer) GetFriends(ctx context.Context, in *emptypb.Empty) (*pb.GetFriendsResp, error) {
	return nil, nil
}

// SendMessageToGroup 发送群组消息
func (*LogicExtServer) SendMessageToGroup(ctx context.Context, in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	return nil, nil
}

func (*LogicExtServer) CreateGroup(ctx context.Context, in *pb.CreateGroupReq) (*pb.CreateGroupResp, error) {
	return nil, nil
}

func (*LogicExtServer) UpdateGroup(ctx context.Context, in *pb.UpdateGroupReq) (*emptypb.Empty, error) {
	return nil, nil
}

func (*LogicExtServer) GetGroup(ctx context.Context, in *pb.GetGroupReq) (*pb.GetGroupResp, error) {
	return nil, nil
}

func (*LogicExtServer) GetGroups(ctx context.Context, in *emptypb.Empty) (*pb.GetGroupsResp, error) {
	return nil, nil
}

func (s *LogicExtServer) AddGroupMembers(ctx context.Context, in *pb.AddGroupMembersReq) (*pb.AddGroupMembersResp, error) {
	return nil, nil
}

func (*LogicExtServer) UpdateGroupMember(ctx context.Context, in *pb.UpdateGroupMemberReq) (*emptypb.Empty, error) {
	return nil, nil
}

func (*LogicExtServer) DeleteGroupMember(ctx context.Context, in *pb.DeleteGroupMemberReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *LogicExtServer) GetGroupMembers(ctx context.Context, in *pb.GetGroupMembersReq) (*pb.GetGroupMembersResp, error) {
	return nil, nil
}
