package service

import (
	"context"
	"go-im/internal/logic/domain/device"
	"go-im/internal/logic/domain/message/model"
	"go-im/internal/logic/domain/message/repo"
	"go-im/pkg/grpclib"
	"go-im/pkg/logger"
	"go-im/pkg/protocol/pb"
	"go-im/pkg/rpc"
	"go-im/pkg/util"

	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

const MessageLimit = 50 // 最大消息同步数量

const MaxSyncBufLen = 65536

type messageService struct{}

var MessageService = new(messageService)

func (*messageService) Sync(ctx context.Context, userId, seq int64) (*pb.SyncResp, error) {
	messages, hasMore, err := MessageService.ListByUserIdAndSeq(ctx, userId, seq)
	if err != nil {
		return nil, err
	}

	pbMessages := model.MessagesToPB(messages)
	length := len(pbMessages)

	resp := &pb.SyncResp{
		Messages: pbMessages,
		HasMore:  hasMore,
	}
	bytes, err := proto.Marshal(resp)
	if err != nil {
		return nil, err
	}

	// 如果字节数组大于一个包的长度，则需要减少字节数组
	for len(bytes) > MaxSyncBufLen {
		length = length * 2 / 3
		resp = &pb.SyncResp{
			Messages: pbMessages[:length],
			HasMore:  hasMore,
		}
		bytes, err = proto.Marshal(resp)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (*messageService) ListByUserIdAndSeq(ctx context.Context, userId, seq int64) ([]model.Message, bool, error) {
	var err error
	if seq == 0 {
		seq, err = DeviceAckService.GetMaxByUserId(ctx, userId)
		if err != nil {
			return nil, false, err
		}
	}
	return repo.MessageRepo.ListBySeq(userId, seq, MessageLimit)
}

// SendToDevice 发送消息给设备
func (*messageService) SendToDevice(ctx context.Context, device *pb.Device, message *pb.Message) error {
	_, err := rpc.GetConnectIntClient().DeliverMessage(ctx, &pb.DeliverMessageReq{
		DeviceId: device.DeviceId,
		Message:  message,
	})
	if err != nil {
		logger.Logger.Error("SendToDevice", zap.Error(err))
	}

	return nil
}

// SendToUser 发送消息给用户
func (*messageService) SendToUser(ctx context.Context, fromDeviceID, toUserID int64, message *pb.Message, isPersist bool) (int64, error) {
	logger.Logger.Debug("SendToUser",
		zap.Int64("request_id", grpclib.GetCtxRequestId(ctx)),
		zap.Int64("to_user_id", toUserID))
	var (
		seq int64 = 0
		err error
	)

	if isPersist {
		seq, err = SeqService.GetUserNext(ctx, toUserID)
		if err != nil {
			return 0, err
		}
		message.Seq = seq

		selfMessage := model.Message{
			UserId:    toUserID,
			RequestId: grpclib.GetCtxRequestId(ctx),
			Code:      message.Code,
			Content:   message.Content,
			Seq:       seq,
			SendTime:  util.UnunixMilliTime(message.SendTime),
			Status:    int32(pb.MessageStatus_MS_NORMAL),
		}
		err = repo.MessageRepo.Save(selfMessage)
		if err != nil {
			logger.Sugar.Error(err)
			return 0, err
		}
	}

	// 查询在线设备, 这里使用了redis
	devices, err := device.App.ListOnlineByUserId(ctx, toUserID)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}

	for i := range devices {
		// 消息不需要推送给发送消息的设备, 这里针对推送给发送者的其他设备的情况
		if devices[i].DeviceId == fromDeviceID {
			continue
		}

		err = MessageService.SendToDevice(ctx, devices[i], message)
		if err != nil {
			logger.Sugar.Error(err, zap.Any("SendToUser error", devices[i]), zap.Error(err))
		}
	}
	return seq, nil
}
