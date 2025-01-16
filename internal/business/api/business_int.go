package api

import (
	"context"
	app2 "go-im/internal/business/domain/user/app"
	"go-im/pkg/protocol/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type BusinessIntServer struct {
	pb.UnsafeBusinessIntServer
}

func (*BusinessIntServer) Auth(ctx context.Context, req *pb.AuthReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, app2.AuthApp.Auth(ctx, req.UserId, req.DeviceId, req.Token)
}

func (*BusinessIntServer) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	user, err := app2.UserApp.Get(ctx, req.UserId)
	return &pb.GetUserResp{User: user}, err
}

func (*BusinessIntServer) GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.GetUsersResp, error) {
	var usersIds = make([]int64, 0, len(req.UserIds))
	for k := range req.UserIds {
		usersIds = append(usersIds, k)
	}
	users, err := app2.UserApp.GetByIds(ctx, usersIds)
	return &pb.GetUsersResp{Users: users}, err
}
