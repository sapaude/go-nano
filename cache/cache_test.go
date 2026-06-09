package cache_test

import (
	"context"
	"testing"
	"time"

	"github.com/sapaude/go-nano/cache"
)

type memCache struct{ data map[string][]byte }

func newMemCache() *memCache { return &memCache{data: make(map[string][]byte)} }
func (m *memCache) Get(_ context.Context, key string) ([]byte, error) {
	v, ok := m.data[key]
	if !ok {
		return nil, nil
	}
	return v, nil
}
func (m *memCache) Set(_ context.Context, key string, value []byte, opts ...cache.SetOption) error {
	_ = cache.ApplySetOptions(opts)
	m.data[key] = value
	return nil
}
func (m *memCache) Delete(_ context.Context, key string) error { delete(m.data, key); return nil }
func (m *memCache) Exists(_ context.Context, key string) (bool, error) {
	_, ok := m.data[key]
	return ok, nil
}

func TestCacheInterface(t *testing.T) {
	var _ cache.Cache = newMemCache()
	c := newMemCache()
	_ = c.Set(context.Background(), "k", []byte("v"), cache.WithTTL(time.Minute))
	v, _ := c.Get(context.Background(), "k")
	if string(v) != "v" {
		t.Fatalf("expected v, got %s", string(v))
	}
}
