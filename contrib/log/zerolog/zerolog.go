package zerologadapter

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/sapaude/go-nano/core"
	flog "github.com/sapaude/go-nano/log"
)

var _ flog.Logger = (*Logger)(nil)
var _ core.Component = (*Logger)(nil)

// Logger adapts zerolog to the go-nano log.Logger interface.
type Logger struct {
	zl zerolog.Logger
}

func New(zl zerolog.Logger) *Logger { return &Logger{zl: zl} }

func (l *Logger) Debug(_ context.Context, msg string, fields ...flog.Field) {
	l.event(l.zl.Debug(), fields).Msg(msg)
}
func (l *Logger) Info(_ context.Context, msg string, fields ...flog.Field) {
	l.event(l.zl.Info(), fields).Msg(msg)
}
func (l *Logger) Warn(_ context.Context, msg string, fields ...flog.Field) {
	l.event(l.zl.Warn(), fields).Msg(msg)
}
func (l *Logger) Error(_ context.Context, msg string, fields ...flog.Field) {
	l.event(l.zl.Error(), fields).Msg(msg)
}
func (l *Logger) With(fields ...flog.Field) flog.Logger {
	ctx := l.zl.With()
	for _, f := range fields {
		ctx = ctx.Interface(f.Key, f.Value)
	}
	return &Logger{zl: ctx.Logger()}
}

func (l *Logger) event(e *zerolog.Event, fields []flog.Field) *zerolog.Event {
	for _, f := range fields {
		e = e.Interface(f.Key, f.Value)
	}
	return e
}

func (l *Logger) Name() string                  { return "log.zerolog" }
func (l *Logger) Init(_ context.Context) error  { return nil }
func (l *Logger) Start(_ context.Context) error { return nil }
func (l *Logger) Stop(_ context.Context) error  { return nil }
