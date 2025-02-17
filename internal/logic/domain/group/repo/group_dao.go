package repo

import (
	"errors"
	"go-im/internal/logic/domain/group/entity"
	"go-im/pkg/db"
	"go-im/pkg/gerrors"

	"gorm.io/gorm"
)

type groupDao struct{}

var GroupDao = new(groupDao)

// Get 获取群组信息
func (*groupDao) Get(groupId int64) (*entity.Group, error) {
	var group = entity.Group{Id: groupId}
	err := db.DB.First(&group).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gerrors.WrapError(err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &group, nil
}

// Save 插入一条群组
func (*groupDao) Save(group *entity.Group) error {
	err := db.DB.Save(&group).Error
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}
