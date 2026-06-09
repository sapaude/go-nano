module github.com/sapaude/go-nano/contrib/cache/redis

go 1.23

require (
	github.com/redis/go-redis/v9 v9.5.1
	github.com/sapaude/go-nano/cache v0.0.0
	github.com/sapaude/go-nano/core v0.0.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	golang.org/x/sync v0.7.0 // indirect
)

replace (
	github.com/sapaude/go-nano/cache => ../../../cache
	github.com/sapaude/go-nano/core => ../../../core
)
