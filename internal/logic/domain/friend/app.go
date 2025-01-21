package friend

import (
	"context"
	"go-im/pkg/protocol/pb"
)

type app struct{}

var App = new(app)

func (s *app) List(ctx context.Context, userId int64) ([]*pb.Friend, error) {
	return Service.List(ctx, userId)
}

func (s *app) SendToFriend(ctx context.Context, fromDeviceID, fromUsedID int64, req *pb.SendMessageReq) (int64, error) {
	return Service.SendToFriend(ctx, fromDeviceID, fromUsedID, req)
}
