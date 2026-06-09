package zerologadapter_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/rs/zerolog"
	zerologadapter "github.com/sapaude/go-nano/contrib/log/zerolog"
	flog "github.com/sapaude/go-nano/log"
)

func TestZerologAdapter(t *testing.T) {
	var buf bytes.Buffer
	zl := zerolog.New(&buf)
	l := zerologadapter.New(zl)

	l.Info(context.Background(), "test message", flog.Field{Key: "k", Value: "v"})
	if buf.Len() == 0 {
		t.Fatal("expected log output")
	}
}

func TestZerologComponent(t *testing.T) {
	l := zerologadapter.New(zerolog.Nop())
	if l.Name() != "log.zerolog" {
		t.Fatalf("unexpected name: %s", l.Name())
	}
	ctx := context.Background()
	if err := l.Init(ctx); err != nil {
		t.Fatal(err)
	}
}
