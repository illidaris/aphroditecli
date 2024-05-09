package log

import (
	"context"
	"fmt"

	"github.com/illidaris/aphrodite/pkg/logex"
	iLog "github.com/illidaris/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var l *zap.Logger
var level iLog.Level = iLog.InfoLevel

func NewLogger() *zap.Logger {
	l = zap.L().WithOptions(zap.AddCallerSkip(2))
	return l
}

func Set(lvl iLog.Level) {
	level = lvl
}

func Debug(ctx context.Context, msg string, args ...interface{}) {
	Log(ctx, fmt.Sprintf(msg, args...), zapcore.DebugLevel)
}

func InfoFocus(ctx context.Context, focus, msg string, args ...interface{}) {
	Log(ctx, fmt.Sprintf(msg, args...), zapcore.InfoLevel, zap.String("focus", focus))
}

func Info(ctx context.Context, msg string, args ...interface{}) {
	Log(ctx, fmt.Sprintf(msg, args...), zapcore.InfoLevel)
}

func Warn(ctx context.Context, msg string, args ...interface{}) {
	Log(ctx, fmt.Sprintf(msg, args...), zapcore.WarnLevel)
}

func WarnFocus(ctx context.Context, focus, msg string, args ...interface{}) {
	Log(ctx, fmt.Sprintf(msg, args...), zapcore.WarnLevel, zap.String("focus", focus))
}

func Error(ctx context.Context, msg string, args ...interface{}) {
	Log(ctx, fmt.Sprintf(msg, args...), zapcore.ErrorLevel)
}

func Log(ctx context.Context, msg string, lvl zapcore.Level, fields ...zapcore.Field) {
	base := logex.FieldsFromCtx(ctx)
	base = append(base, fields...)
	l.Log(lvl, msg, base...)
}
