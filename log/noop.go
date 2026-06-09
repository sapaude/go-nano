package log

import "context"

type noopLogger struct{}

func (n *noopLogger) Debug(_ context.Context, _ string, _ ...Field) {}
func (n *noopLogger) Info(_ context.Context, _ string, _ ...Field)  {}
func (n *noopLogger) Warn(_ context.Context, _ string, _ ...Field)  {}
func (n *noopLogger) Error(_ context.Context, _ string, _ ...Field) {}
func (n *noopLogger) With(_ ...Field) Logger                        { return n }
