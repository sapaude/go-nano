package log

import "context"

// Logger is the logging interface.
type Logger interface {
	Debug(ctx context.Context, msg string, fields ...Field)
	Info(ctx context.Context, msg string, fields ...Field)
	Warn(ctx context.Context, msg string, fields ...Field)
	Error(ctx context.Context, msg string, fields ...Field)
	With(fields ...Field) Logger
}

// Field is a key-value log field.
type Field struct {
	Key   string
	Value any
}

var global Logger = &noopLogger{}

func SetGlobal(l Logger)                                     { global = l }
func Debug(ctx context.Context, msg string, fields ...Field) { global.Debug(ctx, msg, fields...) }
func Info(ctx context.Context, msg string, fields ...Field)  { global.Info(ctx, msg, fields...) }
func Warn(ctx context.Context, msg string, fields ...Field)  { global.Warn(ctx, msg, fields...) }
func Error(ctx context.Context, msg string, fields ...Field) { global.Error(ctx, msg, fields...) }
