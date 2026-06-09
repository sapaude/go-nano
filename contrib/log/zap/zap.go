package zapadapter

import (
	"context"

	"go.uber.org/zap"
	"github.com/sapaude/go-nano/core"
	flog "github.com/sapaude/go-nano/log"
)

var _ flog.Logger = (*Logger)(nil)
var _ core.Component = (*Logger)(nil)

// Logger adapts zap to the go-nano log.Logger interface.
type Logger struct {
	zl *zap.Logger
}

func New(zl *zap.Logger) *Logger { return &Logger{zl: zl} }

func (l *Logger) Debug(_ context.Context, msg string, fields ...flog.Field) {
	l.zl.Debug(msg, toZapFields(fields)...)
}
func (l *Logger) Info(_ context.Context, msg string, fields ...flog.Field) {
	l.zl.Info(msg, toZapFields(fields)...)
}
func (l *Logger) Warn(_ context.Context, msg string, fields ...flog.Field) {
	l.zl.Warn(msg, toZapFields(fields)...)
}
func (l *Logger) Error(_ context.Context, msg string, fields ...flog.Field) {
	l.zl.Error(msg, toZapFields(fields)...)
}
func (l *Logger) With(fields ...flog.Field) flog.Logger {
	return &Logger{zl: l.zl.With(toZapFields(fields)...)}
}

func toZapFields(fields []flog.Field) []zap.Field {
	zf := make([]zap.Field, len(fields))
	for i, f := range fields {
		zf[i] = zap.Any(f.Key, f.Value)
	}
	return zf
}

func (l *Logger) Name() string                  { return "log.zap" }
func (l *Logger) Init(_ context.Context) error  { return nil }
func (l *Logger) Start(_ context.Context) error { return nil }
func (l *Logger) Stop(_ context.Context) error  { return nil }
