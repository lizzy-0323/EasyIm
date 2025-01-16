package repo

import (
	"errors"
	"go-im/internal/business/domain/user/model"
	"go-im/pkg/db"
	"time"

	"gorm.io/gorm"
)

type userDao struct{}

var UserDao = new(userDao)

// Add 插入一条用户信息
func (c *userDao) Add(user model.User) (int64, error) {
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	err := db.DB.Create(&user).Error

	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

// Get 获取用户信息
func (c *userDao) Get(userId int64) (*model.User, error) {
	var user = model.User{Id: userId}
	err := db.DB.First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// Save 保存
func (*userDao) Save(user *model.User) error {
	err := db.DB.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

// GetByPhoneNumber 根据手机号获取用户信息
func (*userDao) GetByPhoneNumber(phoneNumber string) (*model.User, error) {
	var user model.User
	err := db.DB.First(&user, "phone_number = ?", phoneNumber).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

// GetByIds 获取用户信息
func (*userDao) GetByIds(userIds []int64) ([]model.User, error) {
	var users []model.User
	err := db.DB.Find(&users, "id in (?)", userIds).Error
	if err != nil {
		return nil, err
	}
	return users, err
}

// Search 查询用户,这里简单实现，生产环境建议使用ES
func (*userDao) Search(key string) ([]model.User, error) {
	var users []model.User
	key = "%" + key + "%"
	err := db.DB.Where("phone_number like ? or nickname like ?", key, key).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
