package repo

import (
	"errors"
	"go-im/internal/business/domain/user/model"
	"go-im/pkg/db"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	UserKey    = "user:"
	UserExpire = 2 * time.Hour
)

type userCache struct{}

var UserCache = new(userCache)

// 生成随机的过期时间
func genRandomExpire() time.Duration {
	return UserExpire + time.Duration(rand.Intn(3600))*time.Second
}

// Get 获取用户缓存
func (c *userCache) Get(userId int64) (*model.User, error) {
	var user model.User
	err := db.RedisUtil.Get(UserKey+strconv.FormatInt(userId, 10), &user)
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, err
	}
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	return &user, nil
}

// Set 设置用户缓存
func (c *userCache) Set(user model.User) error {
	expireTime := genRandomExpire()
	err := db.RedisUtil.Set(UserKey+strconv.FormatInt(user.Id, 10), user, expireTime)
	if err != nil {
		return err
	}
	return nil
}

// Del 删除用户缓存
func (c *userCache) Del(userId int64) error {
	_, err := db.RedisCli.Del(UserKey + strconv.FormatInt(userId, 10)).Result()
	if err != nil {
		return err
	}
	return nil
}
