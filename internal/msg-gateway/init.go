package msggateway

import (
	"context"

	"go.uber.org/zap"
)

var log *zap.Logger

func init() {
	log, _ = zap.NewProduction()
}

// start msg gateway server
func Start(ctx context.Context, address string, port int) error {
	// Start websocket server
	ws := NewWsServer(address, port)
	netDown := make(chan error)

	if err := ws.Run(netDown); err != nil {
		return err
	}
	return nil
}
