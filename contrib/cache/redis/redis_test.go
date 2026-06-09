package redisadapter_test

import (
	"testing"

	redisadapter "github.com/sapaude/go-nano/contrib/cache/redis"
)

func TestNew(t *testing.T) {
	c := redisadapter.New(redisadapter.Config{Addr: "localhost:6379"})
	if c.Name() != "cache.redis" {
		t.Fatalf("unexpected name: %s", c.Name())
	}
}
