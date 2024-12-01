package msggateway

import (
	"context"
	"net/http"
	"strconv"
	"sync"
	"time"

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
	Port           int
	clientPool     sync.Pool
	registerChan   chan *Client
	unregisterChan chan *Client
}

func NewWsServer(address string, port int) *WsServer {
	return &WsServer{
		address: address,
		Port:    port,
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

func (ws *WsServer) Run(done chan error) error {
	log.Info("Start websocket server at", zap.String("address", ws.address), zap.Int("port", ws.Port))
	// hand websocket connection
	server := http.Server{Addr: ws.address + ":" + strconv.Itoa(ws.Port), Handler: nil}
	http.HandleFunc("/ws", ws.wsHandler)

	// start listening for signal
	// go func() {
	// 	for {
	// 		select {
	// 		case <-ws.registerChan:
	// 			// TODO: add client to pools and check multiple connection
	// 			log.Info("register ws connection")
	// 		case <-ws.unregisterChan:
	// 			// TODO:
	// 			log.Info("unregister ws connection")
	// 		case <-shutdownDone:
	// 			return
	// 		}

	// 	}
	// }()

	// start server
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrAbortHandler {
			done <- err
		}
	}()

	// wait for external signal to stop
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	select {
	case err = <-done:
		sErr := server.Shutdown(ctx)
		if sErr != nil {
			return sErr
		}
		if err != nil {
			return err
		}
		// close(shutdownDone)
		// TODO: add more case
	}

	return nil
}
