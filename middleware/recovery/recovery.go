package recovery

import (
	"context"
	"fmt"
	"runtime/debug"
)

// Handler is a function that processes a request.
type Handler func(ctx context.Context) error

// RecoverFunc is called when a panic is recovered.
type RecoverFunc func(ctx context.Context, p any, stack []byte)

// Middleware wraps a Handler with panic recovery.
func Middleware(recoverFn RecoverFunc) func(Handler) Handler {
	return func(next Handler) Handler {
		return func(ctx context.Context) (err error) {
			defer func() {
				if p := recover(); p != nil {
					stack := debug.Stack()
					if recoverFn != nil {
						recoverFn(ctx, p, stack)
					}
					err = fmt.Errorf("panic: %v", p)
				}
			}()
			return next(ctx)
		}
	}
}
