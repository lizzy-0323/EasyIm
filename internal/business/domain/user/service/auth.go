package service

import (
	"context"
	"errors"
	"go-im/internal/business/domain/user/model"
	"go-im/internal/business/domain/user/repo"
	"go-im/pkg/gerrors"
	"go-im/pkg/protocol/pb"
	"go-im/pkg/rpc"
	"time"
)

type authService struct{}

var AuthService = new(authService)

// SignIn 登录
func (*authService) SignIn(ctx context.Context, phoneNumber, code string, deviceId int64) (bool, int64, string, error) {
	if !Verify(phoneNumber, code) {
		return false, 0, "", errors.New("验证码错误")
	}

	user, err := repo.UserRepo.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return false, 0, "", err
	}
	var isNew = false
	if user == nil {
		user = &model.User{
			PhoneNumber: phoneNumber,
			CreateTime:  time.Now(),
			UpdateTime:  time.Now(),
		}
		err := repo.UserRepo.Save(user)
		if err != nil {
			return false, 0, "", err
		}
		isNew = true
	}

	resp, err := rpc.GetLogicIntClient().GetDevice(ctx, &pb.GetDeviceReq{DeviceId: deviceId})
	if err != nil {
		return false, 0, "", err
	}

	// 测试
	token := "0"

	err = repo.AuthRepo.Set(user.Id, resp.Device.DeviceId, model.Device{
		Type:   resp.Device.Type,
		Token:  token,
		Expire: time.Now().AddDate(0, 3, 0).Unix(),
	})

	if err != nil {
		return false, 0, "", err
	}
	return isNew, user.Id, token, nil
}

func Verify(phoneNumber, code string) bool {
	// TODO: 验证码校验逻辑
	return true
}

func (*authService) Auth(ctx context.Context, userId, deviceId int64, token string) error {
	device, err := repo.AuthRepo.Get(userId, deviceId)
	if err != nil {
		return err
	}

	if device == nil {
		return gerrors.ErrUnauthorized
	}

	if device.Expire < time.Now().Unix() {
		return gerrors.ErrUnauthorized
	}

	if device.Token != token {
		return gerrors.ErrUnauthorized
	}
	return nil
}
