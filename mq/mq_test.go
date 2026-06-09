package mq_test

import (
	"context"
	"testing"

	"github.com/sapaude/go-nano/mq"
)

type mockPublisher struct{ published int }

func (m *mockPublisher) Publish(_ context.Context, _ string, _ []byte) error {
	m.published++
	return nil
}

func TestPublisherInterface(t *testing.T) {
	var _ mq.Publisher = &mockPublisher{}
	p := &mockPublisher{}
	_ = p.Publish(context.Background(), "topic", []byte("hello"))
	if p.published != 1 {
		t.Fatal("expected 1 publish call")
	}
}
