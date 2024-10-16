package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"os"
)

var Logger *zap.SugaredLogger

func init() {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	config := zap.Config{
		Level:         zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:   false,
		Encoding:      "json",
		EncoderConfig: encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	Logger = zap.Must(config.Build()).Sugar()
}
