module github.com/sapaude/go-nano/contrib/log/zap

go 1.23

require (
	github.com/sapaude/go-nano/core v0.0.0
	github.com/sapaude/go-nano/log v0.0.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
)

replace (
	github.com/sapaude/go-nano/core => ../../../core
	github.com/sapaude/go-nano/log => ../../../log
)
