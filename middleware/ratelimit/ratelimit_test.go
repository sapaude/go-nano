package ratelimit_test

import (
	"context"
	"testing"

	"github.com/sapaude/go-nano/middleware/ratelimit"
)

func TestRateLimiter(t *testing.T) {
	l := ratelimit.New(1, 1)
	if !l.Allow() {
		t.Fatal("first allow should succeed")
	}
	if l.Allow() {
		t.Fatal("second allow should be rate limited")
	}
}

func TestRateLimitMiddleware(t *testing.T) {
	l := ratelimit.New(100, 1)
	mw := ratelimit.Middleware(l)
	called := false
	handler := mw(func(_ context.Context) error { called = true; return nil })
	_ = handler(context.Background())
	if !called {
		t.Fatal("handler should have been called")
	}
}
