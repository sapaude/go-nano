package core

import (
	"os"
	"time"
)

type options struct {
	stopTimeout time.Duration
	signals     []os.Signal
}

// Option configures the App.
type Option func(*options)

func WithStopTimeout(d time.Duration) Option {
	return func(o *options) { o.stopTimeout = d }
}

func WithSignals(sigs ...os.Signal) Option {
	return func(o *options) { o.signals = sigs }
}
