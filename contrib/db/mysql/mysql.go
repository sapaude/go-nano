package mysqladapter

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sapaude/go-nano/core"
	"github.com/sapaude/go-nano/db"
)

var _ db.DB = (*DB)(nil)
var _ core.Component = (*DB)(nil)

// Config holds MySQL connection configuration.
type Config struct {
	DSN     string
	MaxOpen int
	MaxIdle int
}

// DB is a MySQL adapter implementing db.DB.
type DB struct {
	cfg Config
	sql *sql.DB
}

func New(cfg Config) *DB { return &DB{cfg: cfg} }

func (d *DB) Name() string { return "db.mysql" }

func (d *DB) Init(_ context.Context) error {
	sqlDB, err := sql.Open("mysql", d.cfg.DSN)
	if err != nil {
		return fmt.Errorf("open mysql: %w", err)
	}
	if d.cfg.MaxOpen > 0 {
		sqlDB.SetMaxOpenConns(d.cfg.MaxOpen)
	}
	if d.cfg.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(d.cfg.MaxIdle)
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
	if err != nil {
		return nil, err
	}
	return rows, nil
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
	return &mysqlTx{tx: tx}, nil
}

type mysqlTx struct {
	tx *sql.Tx
}

func (t *mysqlTx) Query(ctx context.Context, query string, args ...any) (db.Rows, error) {
	rows, err := t.tx.QueryContext(ctx, query, args...)
	return rows, err
}
func (t *mysqlTx) Exec(ctx context.Context, query string, args ...any) error {
	_, err := t.tx.ExecContext(ctx, query, args...)
	return err
}
func (t *mysqlTx) Begin(_ context.Context) (db.Tx, error) {
	return nil, fmt.Errorf("nested transactions not supported")
}
func (t *mysqlTx) Commit(_ context.Context) error   { return t.tx.Commit() }
func (t *mysqlTx) Rollback(_ context.Context) error { return t.tx.Rollback() }
