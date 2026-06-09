package log_test

import (
	"context"
	"testing"

	"github.com/sapaude/go-nano/log"
)

type captureLogger struct {
	last string
}

func (c *captureLogger) Debug(_ context.Context, msg string, _ ...log.Field) { c.last = msg }
func (c *captureLogger) Info(_ context.Context, msg string, _ ...log.Field)  { c.last = msg }
func (c *captureLogger) Warn(_ context.Context, msg string, _ ...log.Field)  { c.last = msg }
func (c *captureLogger) Error(_ context.Context, msg string, _ ...log.Field) { c.last = msg }
func (c *captureLogger) With(_ ...log.Field) log.Logger                      { return c }

func TestSetGlobal(t *testing.T) {
	cl := &captureLogger{}
	log.SetGlobal(cl)
	log.Info(context.Background(), "hello")
	if cl.last != "hello" {
		t.Fatalf("expected 'hello', got '%s'", cl.last)
	}
}

func TestNoopLogger(t *testing.T) {
	// reset to noop proxy
	log.SetGlobal(&noopProxy{})
}

type noopProxy struct{}

func (n *noopProxy) Debug(_ context.Context, _ string, _ ...log.Field) {}
func (n *noopProxy) Info(_ context.Context, _ string, _ ...log.Field)  {}
func (n *noopProxy) Warn(_ context.Context, _ string, _ ...log.Field)  {}
func (n *noopProxy) Error(_ context.Context, _ string, _ ...log.Field) {}
func (n *noopProxy) With(_ ...log.Field) log.Logger                    { return n }
