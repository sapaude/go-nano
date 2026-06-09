package sqliteadapter

import (
	"context"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
	"github.com/sapaude/go-nano/core"
	"github.com/sapaude/go-nano/db"
)

var _ db.DB = (*DB)(nil)
var _ core.Component = (*DB)(nil)

// Config holds SQLite configuration.
type Config struct {
	Path string
}

// DB is a SQLite adapter implementing db.DB.
type DB struct {
	cfg Config
	sql *sql.DB
}

func New(cfg Config) *DB { return &DB{cfg: cfg} }

func (d *DB) Name() string { return "db.sqlite" }

func (d *DB) Init(_ context.Context) error {
	sqlDB, err := sql.Open("sqlite", d.cfg.Path)
	if err != nil {
		return fmt.Errorf("open sqlite: %w", err)
	}
	d.sql = sqlDB
	return nil
}

func (d *DB) Start(_ context.Context) error { return nil }

func (d *DB) Stop(_ context.Context) error {
	if d.sql != nil {
		return d.sql.Close()
	}
	return nil
}

func (d *DB) Query(ctx context.Context, query string, args ...any) (db.Rows, error) {
	rows, err := d.sql.QueryContext(ctx, query, args...)
	return rows, err
}

func (d *DB) Exec(ctx context.Context, query string, args ...any) error {
	_, err := d.sql.ExecContext(ctx, query, args...)
	return err
}

func (d *DB) Begin(ctx context.Context) (db.Tx, error) {
	tx, err := d.sql.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &sqliteTx{tx: tx}, nil
}

type sqliteTx struct {
	tx *sql.Tx
}

func (t *sqliteTx) Query(ctx context.Context, query string, args ...any) (db.Rows, error) {
	return t.tx.QueryContext(ctx, query, args...)
}
func (t *sqliteTx) Exec(ctx context.Context, query string, args ...any) error {
	_, err := t.tx.ExecContext(ctx, query, args...)
	return err
}
func (t *sqliteTx) Begin(_ context.Context) (db.Tx, error) {
	return nil, fmt.Errorf("nested transactions not supported")
}
func (t *sqliteTx) Commit(_ context.Context) error   { return t.tx.Commit() }
func (t *sqliteTx) Rollback(_ context.Context) error { return t.tx.Rollback() }
