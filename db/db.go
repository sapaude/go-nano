package db

import "context"

// Rows is an iterator over query results.
type Rows interface {
	Next() bool
	Scan(dest ...any) error
	Close() error
	Err() error
}

// DB is the database access interface.
type DB interface {
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	Exec(ctx context.Context, sql string, args ...any) error
	Begin(ctx context.Context) (Tx, error)
}
