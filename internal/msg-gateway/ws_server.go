package msggateway

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsServer struct {
	address        string
	clientPool     sync.Pool
	registerChan   chan *Client
	unregisterChan chan *Client
}

func NewWsServer(address string) *WsServer {
	return &WsServer{
		address: address,
		clientPool: sync.Pool{
			New: func() any {
				return new(Client)
			},
		},
	}
}

func (ws *WsServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade http connection to websocket connection
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("upgrade connection failed, err: %v", zap.Error(err))
		return
	}
	client := ws.clientPool.Get().(*Client)
	client.Reset(conn)
	client.ReadMessage()
}

func (ws *WsServer) Run() {
	log.Info("Start websocket server at", zap.String("address", ws.address))

	// hand websocket connection
	server := http.Server{Addr: ws.address, Handler: nil}
	http.HandleFunc("/", ws.wsHandler)

	// start server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
