package db

import "context"

// Tx represents a database transaction.
type Tx interface {
	DB
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
