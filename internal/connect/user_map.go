package connect

import (
	"sync"
)

var UserMap = sync.Map{}

// SetConn 存储
func SetConn(deviceId int64, client *Client) {
	UserMap.Store(deviceId, client)
}

// GetConn 获取
func GetConn(deviceId int64) *Client {
	value, ok := UserMap.Load(deviceId)
	if ok {
		return value.(*Client)
	}
	return nil
}

// DeleteConn 删除
func DeleteConn(deviceId int64) {
	UserMap.Delete(deviceId)
}
