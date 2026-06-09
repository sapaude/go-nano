package kafkaadapter_test

import (
	"testing"

	kafkaadapter "github.com/sapaude/go-nano/contrib/mq/kafka"
)

func TestNew(t *testing.T) {
	p := kafkaadapter.New(kafkaadapter.Config{Brokers: []string{"localhost:9092"}})
	if p.Name() != "mq.kafka" {
		t.Fatalf("unexpected name: %s", p.Name())
	}
}
