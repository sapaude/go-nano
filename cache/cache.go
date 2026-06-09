package cache

import (
	"context"
	"time"
)

// Cache is the key-value cache interface.
type Cache interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value []byte, opts ...SetOption) error
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
}

// SetOption configures a Set operation.
type SetOption func(*SetOptions)

// SetOptions holds the configuration for a Set operation.
type SetOptions struct {
	TTL time.Duration
}

// WithTTL sets the TTL for a cache entry.
func WithTTL(d time.Duration) SetOption {
	return func(o *SetOptions) { o.TTL = d }
}

// ApplySetOptions applies the given options and returns the result.
func ApplySetOptions(opts []SetOption) SetOptions {
	o := SetOptions{}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
