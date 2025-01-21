package service

import (
	"context"
	"go-im/internal/logic/domain/message/model"
	"go-im/internal/logic/domain/message/repo"
	"go-im/pkg/protocol/pb"

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
	// 如果seq不等于0，同步序列号大于seq的信息，否则同步所有信息
	return repo.MessageRepo.ListBySeq(userId, seq, MessageLimit)
}
