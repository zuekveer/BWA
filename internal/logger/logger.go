package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger alias for zap.SugaredLogger to allow easy replacement in the future.
type Logger = *zap.SugaredLogger

// New creates a new zap.SugaredLogger instance.
func New() (Logger, error) {

	l := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				TimeKey:        "ts",
				LevelKey:       "level",
				NameKey:        "logger",
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				CallerKey:      "caller",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(zapcore.InfoLevel),
		),
		zap.AddCaller(),
	)

	zap.ReplaceGlobals(l)
	return l.Sugar(), nil
}
