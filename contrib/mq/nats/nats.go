package natsadapter

import (
	"context"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/sapaude/go-nano/core"
	"github.com/sapaude/go-nano/mq"
)

var _ mq.Publisher = (*Client)(nil)
var _ mq.Subscriber = (*Client)(nil)
var _ core.Component = (*Client)(nil)

// Config holds NATS connection configuration.
type Config struct {
	URL string
}

// Client is a NATS adapter implementing mq.Publisher and mq.Subscriber.
type Client struct {
	cfg  Config
	conn *nats.Conn
}

func New(cfg Config) *Client { return &Client{cfg: cfg} }

func (c *Client) Name() string { return "mq.nats" }

func (c *Client) Init(_ context.Context) error {
	conn, err := nats.Connect(c.cfg.URL)
	if err != nil {
		return fmt.Errorf("connect nats: %w", err)
	}
	c.conn = conn
	return nil
}

func (c *Client) Start(_ context.Context) error { return nil }

func (c *Client) Stop(_ context.Context) error {
	if c.conn != nil {
		c.conn.Close()
	}
	return nil
}

func (c *Client) Publish(_ context.Context, topic string, payload []byte) error {
	return c.conn.Publish(topic, payload)
}

func (c *Client) Subscribe(ctx context.Context, topic string, handler func(context.Context, mq.Message) error) error {
	_, err := c.conn.Subscribe(topic, func(msg *nats.Msg) {
		_ = handler(ctx, mq.Message{Topic: msg.Subject, Payload: msg.Data})
	})
	return err
}
