package main

import (
	"github.com/circutor/ps-go-client/pkg/logger"
	powerstudioapi "github.com/circutor/ps-go-client/pkg/powerStudioAPI"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Example config logger with zap logger.
	cfg := zap.Config{
		Encoding:    "console",
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
		},
	}

	zapLogger, _ := cfg.Build()
	sugaredLogger := zapLogger.Sugar()

	infoLogger := logger.Func(sugaredLogger.Info)

	newLogger := logger.NewLogAdapter(infoLogger)

	ps := powerstudioapi.NewPowerStudio("", "", "", newLogger)
}
