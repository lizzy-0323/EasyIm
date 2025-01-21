package friend

import (
	"errors"
	"go-im/pkg/db"
	"go-im/pkg/gerrors"

	"gorm.io/gorm"
)

type repo struct{}

var Repo = new(repo)

func (*repo) Get(userId, friendId int64) (*Friend, error) {
	friend := Friend{}
	err := db.DB.First(&friend, "user_id = ? AND friend_id = ?", userId, friendId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &friend, nil
}

func (*repo) Save(friend *Friend) error {
	return gerrors.WrapError(db.DB.Save(friend).Error)
}

// List 获取好友列表
func (*repo) List(userId int64, status int) ([]Friend, error) {
	var friends []Friend
	err := db.DB.Where("user_id = ? and status = ?", userId, status).Find(&friends).Error
	return friends, gerrors.WrapError(err)
}
