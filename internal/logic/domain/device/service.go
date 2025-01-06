package device

import (
	"context"
	"go-im/pkg/protocol/pb"
	"go-im/pkg/rpc"
)

type service struct{}

var Service = new(service)

// SignIn 长连接登录
func (*service) SignIn(ctx context.Context, userId, deviceId int64, token string, connAddr string, clientAddr string) error {
	_, err := rpc.GetBusinessIntClient().Auth(ctx, &pb.AuthReq{UserId: userId, DeviceId: deviceId, Token: token})
	if err != nil {
		return err
	}

	// 标记用户在设备上登录
	device, err := Repo.Get(deviceId)
	if err != nil {
		return err
	}
	if device == nil {
		return nil
	}

	device.Online(userId, connAddr, clientAddr)

	err = Repo.Save(device)
	if err != nil {
		return err
	}
	return nil
}
