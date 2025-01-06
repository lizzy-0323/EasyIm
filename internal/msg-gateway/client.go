package msggateway

import (
	"context"
	"go-im/config"
	"go-im/pkg/protocol/pb"
	"go-im/pkg/rpc"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Client struct {
	DeviceId int64
	UserId   int64
	conn     *websocket.Conn
	m        sync.Mutex
}

func (c *Client) GetAddr() string {
	return c.conn.RemoteAddr().String()
}

func (c *Client) HandleSignIn(input *pb.Input) {
	var signIn pb.SignInInput
	err := proto.Unmarshal(input.Data, &signIn)
	if err != nil {
		log.Sugar().Errorf("unmarshal signIn failed, err: %v", err)
		return
	}

	_, err = rpc.GetLogicIntClient().ConnSignIn(context.Background(), &pb.ConnSignInReq{
		UserId:     signIn.UserId,
		DeviceId:   signIn.DeviceId,
		Token:      signIn.Token,
		ConnAddr:   config.Config.ConnectLocalAddr,
		ClientAddr: c.GetAddr(),
	})

	c.Send(pb.PackageType_PT_SIGN_IN, input.RequestId, nil, err)
	if err != nil {
		return
	}

	c.UserId = signIn.UserId
	c.DeviceId = signIn.DeviceId
	SetConn(signIn.DeviceId, c)
}

func (c *Client) Close() {
	if c.DeviceId != 0 {
		DeleteConn(c.DeviceId)
	}

	if c.DeviceId != 0 {
		_, _ = rpc.GetLogicIntClient().Offline(context.TODO(), &pb.OfflineReq{
			UserId:     c.UserId,
			DeviceId:   c.DeviceId,
			ClientAddr: c.GetAddr(),
		})
	}
	// close websocket connection
	c.conn.Close()
}

func (c *Client) HandleHeartbeat(input *pb.Input) {
	c.Send(pb.PackageType_PT_HEARTBEAT, input.RequestId, nil, nil)

	log.Sugar().Infow("heartbeat", "UserId", c.UserId, "DeviceId", c.DeviceId)
}

// MessageAck 消息收到回执
func (c *Client) MessageAck(input *pb.Input) {
	var messageAck pb.MessageACK
	err := proto.Unmarshal(input.Data, &messageAck)
	if err != nil {
		log.Sugar().Error(err)
		return
	}

	_, err = rpc.GetLogicIntClient().MessageACK(context.TODO(), &pb.MessageACKReq{
		UserId:      c.UserId,
		DeviceId:    c.DeviceId,
		DeviceAck:   messageAck.DeviceAck,
		ReceiveTime: messageAck.ReceiveTime,
	})
}

// Send
func (c *Client) Send(pt pb.PackageType, requestId int64, message proto.Message, err error) {
	var output = pb.Output{
		Type:      pt,
		RequestId: requestId,
	}

	if err != nil {
		status, _ := status.FromError(err)
		output.Code = int32(status.Code())
		output.Message = status.Message()
	}

	if message != nil {
		data, err := proto.Marshal(message)
		if err != nil {
			log.Sugar().Error(err)
			return
		}
		output.Data = data
	}

	outputBytes, err := proto.Marshal(&output)
	if err != nil {
		log.Sugar().Error(err)
		return
	}

	err = c.Write(outputBytes)
	if err != nil {
		log.Sugar().Error(err)
		c.Close()
		return
	}
}

func (c *Client) Write(msg []byte) error {
	// write to websocket
	c.m.Lock()
	defer c.m.Unlock()
	err := c.conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.BinaryMessage, msg)
}

// HandleMessage handle the message from websocket
func (c *Client) HandleMessage(bytes []byte) {
	var input = new(pb.Input)
	err := proto.Unmarshal(bytes, input)
	if err != nil {
		log.Error("unmarshal message failed", zap.Error(err), zap.Int("len", len(bytes)))
		return
	}
	log.Debug("Handle message", zap.Any("input", input))

	switch input.Type {
	case pb.PackageType_PT_SIGN_IN:
		c.HandleSignIn(input)
	case pb.PackageType_PT_HEARTBEAT:
		c.HandleHeartbeat(input)
	case pb.PackageType_PT_MESSAGE:
		c.MessageAck(input)
	default:
		log.Error("unknown message type", zap.Int32("type", int32(input.Type)))
	}
}

func (c *Client) Reset(conn *websocket.Conn) {
	c.conn = conn
	c.m = sync.Mutex{}
	c.DeviceId = 0
	c.UserId = 0
}

// read message
func (c *Client) ReadMessage() {
	// recover from panic
	defer func() {
		r := recover()
		if r != nil {
			log.Error("panic", zap.Any("recover", r))
		}
	}()

	// handle websocket connection
	for {
		err := c.conn.SetReadDeadline(time.Now().Add(12 * time.Minute))
		if err != nil {
			log.Error("set read deadline failed", zap.Error(err))
			break
		}
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Error("read message failed", zap.Error(err))
			return
		}

		c.HandleMessage(msg)
	}
}
