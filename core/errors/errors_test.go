package errors_test

import (
	stderrors "errors"
	"testing"

	"github.com/sapaude/go-nano/core/errors"
)

func TestNew(t *testing.T) {
	err := errors.New("test.code", "test message")
	if err.Code != "test.code" {
		t.Fatalf("expected code test.code, got %s", err.Code)
	}
	if err.Error() != "[test.code] test message" {
		t.Fatalf("unexpected error string: %s", err.Error())
	}
}

func TestWrap(t *testing.T) {
	cause := stderrors.New("underlying")
	err := errors.Wrap(cause, "wrap.code", "wrapped")
	if !stderrors.Is(err, cause) {
		t.Fatal("errors.Is should work through Wrap")
	}
	if errors.Code(err) != "wrap.code" {
		t.Fatalf("unexpected code: %s", errors.Code(err))
	}
}

func TestCode_unknown(t *testing.T) {
	err := stderrors.New("plain error")
	if errors.Code(err) != "unknown" {
		t.Fatalf("expected unknown, got %s", errors.Code(err))
	}
}
