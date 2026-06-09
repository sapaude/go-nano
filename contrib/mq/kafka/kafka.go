package kafkaadapter

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/sapaude/go-nano/core"
	"github.com/sapaude/go-nano/mq"
)

var _ mq.Publisher = (*Publisher)(nil)
var _ core.Component = (*Publisher)(nil)

// Config holds Kafka producer configuration.
type Config struct {
	Brokers []string
	Topic   string
}

// Publisher is a Kafka adapter implementing mq.Publisher.
type Publisher struct {
	cfg      Config
	producer sarama.SyncProducer
}

func New(cfg Config) *Publisher { return &Publisher{cfg: cfg} }

func (p *Publisher) Name() string { return "mq.kafka" }

func (p *Publisher) Init(_ context.Context) error {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(p.cfg.Brokers, cfg)
	if err != nil {
		return fmt.Errorf("create kafka producer: %w", err)
	}
	p.producer = producer
	return nil
}

func (p *Publisher) Start(_ context.Context) error { return nil }

func (p *Publisher) Stop(_ context.Context) error {
	if p.producer != nil {
		return p.producer.Close()
	}
	return nil
}

func (p *Publisher) Publish(_ context.Context, topic string, payload []byte) error {
	_, _, err := p.producer.SendMessage(&sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(payload),
	})
	return err
}
