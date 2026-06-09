package recovery_test

import (
	"context"
	"testing"

	"github.com/sapaude/go-nano/middleware/recovery"
)

func TestRecovery(t *testing.T) {
	var caught any
	mw := recovery.Middleware(func(_ context.Context, p any, _ []byte) { caught = p })
	handler := mw(func(_ context.Context) error {
		panic("test panic")
	})
	err := handler(context.Background())
	if err == nil {
		t.Fatal("expected error from panic")
	}
	if caught != "test panic" {
		t.Fatalf("expected caught panic 'test panic', got %v", caught)
	}
}
