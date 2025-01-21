package proxy

import (
	"context"
	"go-im/pkg/protocol/pb"
)

type deviceProxy interface {
	ListOnlineByUserId(ctx context.Context, userId int64) ([]*pb.Device, error)
}

var DeviceProxy deviceProxy
