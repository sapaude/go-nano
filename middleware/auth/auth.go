package auth

import (
	"context"
	"errors"
)

var ErrUnauthorized = errors.New("unauthorized")

// TokenVerifier verifies a token string and returns the subject.
type TokenVerifier interface {
	Verify(ctx context.Context, token string) (subject string, err error)
}

type contextKey struct{}

// TokenFromContext extracts the subject set by the auth middleware.
func TokenFromContext(ctx context.Context) (string, bool) {
	s, ok := ctx.Value(contextKey{}).(string)
	return s, ok
}

type Handler func(ctx context.Context, token string) error

func Middleware(v TokenVerifier) func(Handler) Handler {
	return func(next Handler) Handler {
		return func(ctx context.Context, token string) error {
			subject, err := v.Verify(ctx, token)
			if err != nil {
				return ErrUnauthorized
			}
			ctx = context.WithValue(ctx, contextKey{}, subject)
			return next(ctx, token)
		}
	}
}
