module github.com/sapaude/go-nano/contrib/db/mysql

go 1.23

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/sapaude/go-nano/core v0.0.0
	github.com/sapaude/go-nano/db v0.0.0
)

require golang.org/x/sync v0.7.0 // indirect

replace (
	github.com/sapaude/go-nano/core => ../../../core
	github.com/sapaude/go-nano/db => ../../../db
)
