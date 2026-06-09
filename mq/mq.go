package mq

import "context"

// Message represents a message received from a topic.
type Message struct {
	Topic   string
	Payload []byte
	Headers map[string]string
}

// Publisher publishes messages to a topic.
type Publisher interface {
	Publish(ctx context.Context, topic string, payload []byte) error
}

// Subscriber subscribes to messages from a topic.
type Subscriber interface {
	Subscribe(ctx context.Context, topic string, handler func(ctx context.Context, msg Message) error) error
}
