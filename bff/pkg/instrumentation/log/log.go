package log

import (
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"

	"go.uber.org/zap"
	zapadapter "logur.dev/adapter/zap"
	"logur.dev/logur"
)

type Logger = logur.Logger
type Fields = logur.Fields
type contextKey int

var WithFields = logur.WithFields

const (
	logCtxKey contextKey = iota
)
const (
	TagError string = "err"
)

func NewLogger(ctx context.Context, cfg *Config, opts ...zap.Option) (logur.Logger, error) {
	opts = append(opts, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	if !strings.Contains(strings.ToLower(cfg.Environment), "dev") {
		l, err := zap.NewProduction(opts...)
		if err != nil {
			return nil, err
		}
		defer func() {
			if syncErr := l.Sync(); syncErr != nil {
				fmt.Print(syncErr)
			}
		}()
		return zapadapter.New(l), nil
	}
	var core zapcore.Core
	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	opts = append(opts, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core = zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)
	logger := zap.New(core)
	logger.WithOptions(opts...)
	defer func() {
		if syncErr := logger.Sync(); syncErr != nil {
			fmt.Print(syncErr)
		}
	}()
	return zapadapter.New(logger), nil
}

func WithLogger(ctx context.Context, logger logur.Logger) context.Context {
	return context.WithValue(ctx, logCtxKey, logger)
}

// GetLogger tries to get the logger out of context.
// If one cannot be found, it will wrap the global zap logger with the interface.
func GetLogger(ctx context.Context) logur.Logger {

	if l, ok := ctx.Value(logCtxKey).(logur.Logger); ok {
		return l
	}
	l, _ := zap.NewDevelopment()
	return zapadapter.New(l)
}
