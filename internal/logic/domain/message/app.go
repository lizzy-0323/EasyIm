package message

import (
	"context"
	"go-im/internal/logic/domain/message/service"
	"go-im/pkg/protocol/pb"
)

type app struct{}

var App = new(app)

// Sync 同步消息
func (*app) Sync(ctx context.Context, userId, seq int64) (*pb.SyncResp, error) {
	return service.MessageService.Sync(ctx, userId, seq)
}

// MessageAck 收到消息回执
func (*app) MessageAck(ctx context.Context, userId, deviceId, ack int64) error {
	return service.DeviceAckService.Update(ctx, userId, deviceId, ack)
}
