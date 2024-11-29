package msggateway

import (
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const (
	MessageText = iota + 1
	MessageClose
	MessagePing
	MessagePong
)

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		conn: conn,
		m:    sync.Mutex{},
	}
}

type Client struct {
	conn *websocket.Conn
	m    sync.Mutex
}

func (c *Client) HandleMessage(msg []byte, msgType int) {
	// handle message
	switch msgType {
	case MessageText:
		// handle text message
	case MessageClose:
		return
	case MessagePing:
		// handle ping message
	case MessagePong:
		// handle pong message
	default:
	}
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
		err := c.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		if err != nil {
			log.Error("set read deadline failed", zap.Error(err))
			break
		}
		msgType, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Error("read message failed", zap.Error(err))
			return
		}
		log.Info("recv message: %s", zap.String("msg", string(msg)))
		c.HandleMessage(msg, msgType)
	}
}
