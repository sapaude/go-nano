package core

import "context"

// Component is the interface all framework components must implement.
type Component interface {
	Name() string
	Init(ctx context.Context) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

// Healthchecker is an optional interface for components that support health checks.
type Healthchecker interface {
	Health(ctx context.Context) error
}
