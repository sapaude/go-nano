package redisadapter

import (
	"context"

	goredis "github.com/redis/go-redis/v9"
	"github.com/sapaude/go-nano/cache"
	"github.com/sapaude/go-nano/core"
)

var _ cache.Cache = (*Cache)(nil)
var _ core.Component = (*Cache)(nil)

// Config holds Redis connection configuration.
type Config struct {
	Addr     string
	Password string
	DB       int
}

// Cache is a Redis adapter implementing cache.Cache.
type Cache struct {
	cfg    Config
	client *goredis.Client
}

func New(cfg Config) *Cache { return &Cache{cfg: cfg} }

func (c *Cache) Name() string { return "cache.redis" }

func (c *Cache) Init(_ context.Context) error {
	c.client = goredis.NewClient(&goredis.Options{
		Addr:     c.cfg.Addr,
		Password: c.cfg.Password,
		DB:       c.cfg.DB,
	})
	return nil
}

func (c *Cache) Start(_ context.Context) error { return nil }

func (c *Cache) Stop(_ context.Context) error {
	if c.client != nil {
		return c.client.Close()
	}
	return nil
}

func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	val, err := c.client.Get(ctx, key).Bytes()
	if err == goredis.Nil {
		return nil, nil
	}
	return val, err
}

func (c *Cache) Set(ctx context.Context, key string, value []byte, opts ...cache.SetOption) error {
	o := cache.ApplySetOptions(opts)
	return c.client.Set(ctx, key, value, o.TTL).Err()
}

func (c *Cache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *Cache) Exists(ctx context.Context, key string) (bool, error) {
	n, err := c.client.Exists(ctx, key).Result()
	return n > 0, err
}
