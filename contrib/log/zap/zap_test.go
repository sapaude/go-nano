package zapadapter_test

import (
	"context"
	"testing"

	"go.uber.org/zap"
	zapadapter "github.com/sapaude/go-nano/contrib/log/zap"
	flog "github.com/sapaude/go-nano/log"
)

func TestZapAdapter(t *testing.T) {
	zl, _ := zap.NewDevelopment()
	l := zapadapter.New(zl)
	l.Info(context.Background(), "test", flog.Field{Key: "k", Value: 1})
	if l.Name() != "log.zap" {
		t.Fatalf("unexpected name: %s", l.Name())
	}
}
