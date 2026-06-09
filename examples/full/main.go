package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sapaude/go-nano/core"
	"github.com/sapaude/go-nano/log"
)

type noopComponent struct{ name string }

func (n *noopComponent) Name() string                    { return n.name }
func (n *noopComponent) Init(_ context.Context) error    { return nil }
func (n *noopComponent) Start(ctx context.Context) error { <-ctx.Done(); return nil }
func (n *noopComponent) Stop(_ context.Context) error    { return nil }

func main() {
	log.Info(context.Background(), "starting go-nano example")

	app := core.New(core.WithStopTimeout(5 * time.Second)).Register(
		&noopComponent{name: "example.component"},
	)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		fmt.Printf("app exited: %v\n", err)
	}
}
