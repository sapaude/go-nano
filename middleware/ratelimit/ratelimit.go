package ratelimit

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrRateLimited = errors.New("rate limited")

// Limiter is a token-bucket rate limiter.
type Limiter struct {
	mu       sync.Mutex
	tokens   float64
	maxToken float64
	rate     float64 // tokens per second
	last     time.Time
}

func New(rate float64, burst int) *Limiter {
	return &Limiter{
		tokens:   float64(burst),
		maxToken: float64(burst),
		rate:     rate,
		last:     time.Now(),
	}
}

func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(l.last).Seconds()
	l.last = now
	l.tokens += elapsed * l.rate
	if l.tokens > l.maxToken {
		l.tokens = l.maxToken
	}
	if l.tokens < 1 {
		return false
	}
	l.tokens--
	return true
}

type Handler func(ctx context.Context) error

func Middleware(l *Limiter) func(Handler) Handler {
	return func(next Handler) Handler {
		return func(ctx context.Context) error {
			if !l.Allow() {
				return ErrRateLimited
			}
			return next(ctx)
		}
	}
}
