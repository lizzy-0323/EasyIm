package repo

import (
	"encoding/json"
	"errors"
	"go-im/internal/business/domain/user/model"
	"go-im/pkg/db"
	"strconv"

	"github.com/go-redis/redis"
)

// 每一个user对应一个auth table
const (
	AuthKey = "auth:"
)

type authCache struct{}

var AuthCache = new(authCache)

func (*authCache) Get(userId, deviceId int64) (*model.Device, error) {
	bytes, err := db.RedisCli.HGet(AuthKey+strconv.FormatInt(userId, 10), strconv.FormatInt(deviceId, 10)).Bytes()
	if err != nil && errors.Is(err, redis.Nil) {
		return nil, err
	}

	if errors.Is(err, redis.Nil) {
		return nil, nil
	}

	var device model.Device
	err = json.Unmarshal(bytes, &device)
	if err != nil {
		return nil, err
	}

	return &device, nil
}

func (*authCache) Set(userId, deviceId int64, device model.Device) error {
	bytes, err := json.Marshal(device)
	if err != nil {
		return err
	}

	_, err = db.RedisCli.HSet(AuthKey+strconv.FormatInt(userId, 10), strconv.FormatInt(deviceId, 10), bytes).Result()
	if err != nil {
		return err
	}
	return nil
}

func (*authCache) GetAll(userId int64) (map[int64]model.Device, error) {
	result, err := db.RedisCli.HGetAll(AuthKey + strconv.FormatInt(userId, 10)).Result()
	if err != nil {
		return nil, err
	}

	var devices = make(map[int64]model.Device, len(result))

	for k, v := range result {
		deviceId, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return nil, err
		}

		var device model.Device
		err = json.Unmarshal([]byte(v), &device)
		if err != nil {
			return nil, err
		}
		devices[deviceId] = device
	}
	return devices, nil
}
