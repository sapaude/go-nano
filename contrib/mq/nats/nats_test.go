package natsadapter_test

import (
	"testing"

	natsclient "github.com/nats-io/nats.go"
	natsadapter "github.com/sapaude/go-nano/contrib/mq/nats"
)

func TestNew(t *testing.T) {
	c := natsadapter.New(natsadapter.Config{URL: natsclient.DefaultURL})
	if c.Name() != "mq.nats" {
		t.Fatalf("unexpected name: %s", c.Name())
	}
}
