package http

import (
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewFileLogger() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   "request.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     1,
		Compress:   true,
	}
}

func NewLogger(fileLogger *lumberjack.Logger) zerolog.Logger {
	return zerolog.New(fileLogger)
}

var LoggerModule = fx.Module("http/logger",
	fx.Provide(NewFileLogger, NewLogger),
)
