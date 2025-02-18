package connect

import (
	"container/list"
	"go-im/pkg/protocol/pb"
	"sync"
)

var RoomsManager sync.Map

type Room struct {
	RoomId int64      // 房间ID
	Conns  *list.List // 订阅房间消息的连接
	lock   sync.RWMutex
}

func NewRoom(roomId int64) *Room {
	return &Room{
		RoomId: roomId,
		Conns:  list.New(),
	}
}

// SubscribedRoom 订阅房间
func SubscribedRoom(c *Client, roomId int64) {
	if roomId == c.RoomId {
		return
	}

	oldRoomId := c.RoomId
	// 取消订阅
	if oldRoomId != 0 {
		value, ok := RoomsManager.Load(oldRoomId)
		if !ok {
			return
		}
		room := value.(*Room)
		room.Unsubscribe(c)

		if room.Conns.Front() == nil {
			RoomsManager.Delete(oldRoomId)
		}
		return
	}

	// 订阅
	if roomId != 0 {
		value, ok := RoomsManager.Load(roomId)
		var room *Room
		if !ok {
			room = NewRoom(roomId)
			RoomsManager.Store(roomId, room)
		} else {
			room = value.(*Room)
		}
		room.Subscribe(c)
		return
	}
}

// Subscribe 订阅房间
func (r *Room) Subscribe(c *Client) {
	r.lock.Lock()
	defer r.lock.Unlock()

	c.Element = r.Conns.PushBack(c)
	c.RoomId = r.RoomId
}

// Unsubscribe 取消订阅
func (r *Room) Unsubscribe(c *Client) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.Conns.Remove(c.Element)
	c.Element = nil
	c.RoomId = 0
}

// Push 推送消息到房间
func (r *Room) Push(message *pb.Message) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	element := r.Conns.Front()
	for {
		conn := element.Value.(*Client)
		conn.Send(pb.PackageType_PT_MESSAGE, 0, message, nil)

		element = element.Next()
		if element == nil {
			break
		}
	}
}
