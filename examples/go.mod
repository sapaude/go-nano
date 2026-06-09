module github.com/sapaude/go-nano/examples

go 1.23

require (
	github.com/sapaude/go-nano/core v0.0.0
	github.com/sapaude/go-nano/log v0.0.0
)

require golang.org/x/sync v0.7.0 // indirect

replace (
	github.com/sapaude/go-nano/core => ../core
	github.com/sapaude/go-nano/log => ../log
)
