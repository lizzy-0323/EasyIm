package device

import (
	"context"
	"go-im/pkg/gerrors"
	"go-im/pkg/protocol/pb"
)

type app struct{}

var App = new(app)

// SignIn 登录
func (*app) SignIn(ctx context.Context, userId, deviceId int64, token string, connAddr string, clientAddr string) error {
	return Service.SignIn(ctx, userId, deviceId, token, connAddr, clientAddr)
}

// GetDevice 获取设备信息
func (*app) GetDevice(ctx context.Context, deviceId int64) (*pb.Device, error) {
	device, err := Repo.Get(deviceId)
	if err != nil {
		return nil, err
	}
	if device == nil {
		return nil, gerrors.ErrDeviceNotExist
	}

	return device.ToProto(), err
}

func (*app) Offline(ctx context.Context, deviceId int64, clientAddr string) error {
	device, err := Repo.Get(deviceId)
	if err != nil {
		return err
	}
	if device == nil {
		return nil
	}

	if device.ClientAddr != clientAddr {
		return nil
	}
	device.Status = DeviceOffLine

	err = Repo.Save(device)
	if err != nil {
		return err
	}
	return nil
}

func (*app) Register(ctx context.Context, in *pb.RegisterDeviceReq) (int64, error) {
	device := Device{
		Type:          in.Type,
		Brand:         in.Brand,
		Model:         in.Model,
		SystemVersion: in.SystemVersion,
		SDKVersion:    in.SdkVersion,
	}

	if !device.IsLegal() {
		return 0, gerrors.ErrBadRequest
	}

	err := Repo.Save(&device)
	if err != nil {
		return 0, err
	}

	return device.Id, nil
}

func (*app) ListOnlineByUserId(ctx context.Context, userId int64) ([]*pb.Device, error) {
	return Service.ListOnlineByUserId(ctx, userId)
}

func (*app) ServerStop(ctx context.Context, connAddr string) error {
	return Service.ServerStop(ctx, connAddr)
}
