package http

import (
	"github.com/nexters/book/config"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewFileLogger(settings *config.Settings) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   settings.App.LOGGER,
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
