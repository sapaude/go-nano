module github.com/sapaude/go-nano/contrib/log/zerolog

go 1.23

require (
	github.com/rs/zerolog v1.32.0
	github.com/sapaude/go-nano/core v0.0.0
	github.com/sapaude/go-nano/log v0.0.0
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
)

replace (
	github.com/sapaude/go-nano/core => ../../../core
	github.com/sapaude/go-nano/log => ../../../log
)
