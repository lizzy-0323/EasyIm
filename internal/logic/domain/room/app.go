package room

import (
	"context"
	"go-im/pkg/protocol/pb"
)

type app struct{}

var App = new(app)

func (s *app) Push(ctx context.Context, req *pb.PushRoomReq) error {
	return Service.Push(ctx, req)
}

func (s *app) SubscribeRoom(ctx context.Context, req *pb.SubscribeRoomReq) error {
	return Service.SubscribeRoom(ctx, req)
}
