package friend

import (
	"context"
	"go-im/pkg/protocol/pb"
	"time"
)

type app struct{}

var App = new(app)

func (s *app) List(ctx context.Context, userId int64) ([]*pb.Friend, error) {
	return Service.List(ctx, userId)
}

func (s *app) SendToFriend(ctx context.Context, fromDeviceID, fromUsedID int64, req *pb.SendMessageReq) (int64, error) {
	return Service.SendToFriend(ctx, fromDeviceID, fromUsedID, req)
}

// AddFriend 添加好友
func (*app) AddFriend(ctx context.Context, userId, friendId int64, remarks, description string) error {
	return Service.AddFriend(ctx, userId, friendId, remarks, description)
}

// AgreeAddFriend 同意添加好友
func (*app) AgreeAddFriend(ctx context.Context, userId, friendId int64, remarks string) error {
	return Service.AgreeAddFriend(ctx, userId, friendId, remarks)
}

// SetFriend 设置好友信息
func (*app) SetFriend(ctx context.Context, userId int64, req *pb.SetFriendReq) error {
	friend, err := Repo.Get(userId, req.FriendId)
	if err != nil {
		return err
	}
	if friend == nil {
		return nil
	}

	friend.Remarks = req.Remarks
	friend.Extra = req.Extra
	friend.UpdateTime = time.Now()

	err = Repo.Save(friend)
	if err != nil {
		return err
	}
	return nil
}
