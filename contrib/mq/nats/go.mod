module github.com/sapaude/go-nano/contrib/mq/nats

go 1.23

require (
	github.com/nats-io/nats.go v1.34.0
	github.com/sapaude/go-nano/core v0.0.0
	github.com/sapaude/go-nano/mq v0.0.0
)

require (
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/nats-io/nkeys v0.4.7 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.19.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
)

replace (
	github.com/sapaude/go-nano/core => ../../../core
	github.com/sapaude/go-nano/mq => ../../../mq
)
