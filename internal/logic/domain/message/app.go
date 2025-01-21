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

// SendToUser 发送消息给用户
func (*app) SendToUser(ctx context.Context, fromDeviceID, toUserID int64, message *pb.Message, isPersist bool) (int64, error) {
	return service.MessageService.SendToUser(ctx, fromDeviceID, toUserID, message, isPersist)
}
