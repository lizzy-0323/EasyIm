package logger

import "go.uber.org/zap"

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func init() {
	Logger = zap.NewExample()
	Sugar = Logger.Sugar()
}
