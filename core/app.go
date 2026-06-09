package core

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

// App orchestrates the lifecycle of registered components.
type App struct {
	components []Component
	opts       options
}

func New(opts ...Option) *App {
	a := &App{
		opts: options{
			stopTimeout: 30 * time.Second,
			signals:     []os.Signal{syscall.SIGTERM, syscall.SIGINT},
		},
	}
	for _, o := range opts {
		o(&a.opts)
	}
	return a
}

func (a *App) Register(components ...Component) *App {
	a.components = append(a.components, components...)
	return a
}

func (a *App) Run(ctx context.Context) error {
	for _, c := range a.components {
		if err := c.Init(ctx); err != nil {
			return fmt.Errorf("init %s: %w", c.Name(), err)
		}
	}

	eg, egCtx := errgroup.WithContext(ctx)
	for _, c := range a.components {
		c := c
		eg.Go(func() error { return c.Start(egCtx) })
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, a.opts.signals...)
	select {
	case <-quit:
	case <-egCtx.Done():
	}

	stopCtx, cancel := context.WithTimeout(context.Background(), a.opts.stopTimeout)
	defer cancel()
	for i := len(a.components) - 1; i >= 0; i-- {
		if err := a.components[i].Stop(stopCtx); err != nil {
			fmt.Fprintf(os.Stderr, "stop %s: %v\n", a.components[i].Name(), err)
		}
	}
	return eg.Wait()
}
