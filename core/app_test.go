package core_test

import (
	"context"
	"testing"
	"time"

	"github.com/sapaude/go-nano/core"
)

type mockComponent struct {
	name    string
	started bool
	stopped bool
}

func (m *mockComponent) Name() string                    { return m.name }
func (m *mockComponent) Init(ctx context.Context) error  { return nil }
func (m *mockComponent) Start(ctx context.Context) error { m.started = true; <-ctx.Done(); return nil }
func (m *mockComponent) Stop(ctx context.Context) error  { m.stopped = true; return nil }

func TestApp_RegisterAndRun(t *testing.T) {
	c := &mockComponent{name: "test"}
	app := core.New(core.WithStopTimeout(1 * time.Second)).Register(c)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		// context cancellation is expected, not an error to fail on
		_ = err
	}
	if !c.started {
		t.Fatal("component should have been started")
	}
}
