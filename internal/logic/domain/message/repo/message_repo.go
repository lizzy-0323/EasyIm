package repo

import (
	"go-im/internal/logic/domain/message/model"
	"go-im/pkg/db"
	"go-im/pkg/gerrors"
)

type messageRepo struct{}

var MessageRepo = new(messageRepo)

const MESSAGE_TABLE = "message"

// Save 插入一条消息
func (d *messageRepo) Save(message model.Message) error {
	err := db.DB.Table(MESSAGE_TABLE).Create(&message).Error
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

// ListBySeq 根据类型和id查询大于序号大于seq的消息
func (d *messageRepo) ListBySeq(userId, seq, limit int64) ([]model.Message, bool, error) {
	DB := db.DB.Table(MESSAGE_TABLE).Where("user_id = ? and seq > ?", userId, seq)
	var count int64
	err := DB.Count(&count).Error
	if err != nil {
		return nil, false, gerrors.WrapError(err)
	}
	if count == 0 {
		return nil, false, nil
	}

	var messages []model.Message
	err = DB.Limit(int(limit)).Find(&messages).Error
	if err != nil {
		return nil, false, gerrors.WrapError(err)
	}
	return messages, count > limit, nil
}
