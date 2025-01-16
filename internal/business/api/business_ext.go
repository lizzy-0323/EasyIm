package api

import (
	"context"
	app2 "go-im/internal/business/domain/user/app"
	"go-im/pkg/protocol/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type BusinessExtServer struct {
	pb.UnsafeBusinessExtServer
}

func (s *BusinessExtServer) SignIn(ctx context.Context, req *pb.SignInReq) (*pb.SignInResp, error) {
	isNew, userId, token, err := app2.AuthApp.SignIn(ctx, req.PhoneNumber, req.Code, req.DeviceId)
	if err != nil {
		return nil, err
	}
	return &pb.SignInResp{
		IsNew:  isNew,
		UserId: userId,
		Token:  token,
	}, nil
}

func (s *BusinessExtServer) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	return nil, nil
}

func (s *BusinessExtServer) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *BusinessExtServer) SearchUser(ctx context.Context, req *pb.SearchUserReq) (*pb.SearchUserResp, error) {
	return nil, nil
}
