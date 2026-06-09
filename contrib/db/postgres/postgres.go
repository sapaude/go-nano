package pgadapter

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sapaude/go-nano/core"
	"github.com/sapaude/go-nano/db"
)

var _ db.DB = (*DB)(nil)
var _ core.Component = (*DB)(nil)

// Config holds PostgreSQL connection configuration.
type Config struct {
	DSN      string
	MaxConns int32
}

// DB is a PostgreSQL adapter implementing db.DB.
type DB struct {
	cfg  Config
	pool *pgxpool.Pool
}

func New(cfg Config) *DB { return &DB{cfg: cfg} }

func (d *DB) Name() string { return "db.postgres" }

func (d *DB) Init(ctx context.Context) error {
	config, err := pgxpool.ParseConfig(d.cfg.DSN)
	if err != nil {
		return fmt.Errorf("parse dsn: %w", err)
	}
	if d.cfg.MaxConns > 0 {
		config.MaxConns = d.cfg.MaxConns
	}
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return fmt.Errorf("create pool: %w", err)
	}
	d.pool = pool
	return nil
}

func (d *DB) Start(_ context.Context) error { return nil }

func (d *DB) Stop(_ context.Context) error {
	if d.pool != nil {
		d.pool.Close()
	}
	return nil
}

func (d *DB) Query(ctx context.Context, sql string, args ...any) (db.Rows, error) {
	rows, err := d.pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return &pgRows{rows: rows}, nil
}

func (d *DB) Exec(ctx context.Context, sql string, args ...any) error {
	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}

func (d *DB) Begin(ctx context.Context) (db.Tx, error) {
	tx, err := d.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &pgTx{tx: tx}, nil
}

type pgRows struct {
	rows pgx.Rows
}

func (r *pgRows) Next() bool             { return r.rows.Next() }
func (r *pgRows) Scan(dest ...any) error { return r.rows.Scan(dest...) }
func (r *pgRows) Close() error           { r.rows.Close(); return nil }
func (r *pgRows) Err() error             { return r.rows.Err() }

type pgTx struct {
	tx pgx.Tx
}

func (t *pgTx) Query(ctx context.Context, sql string, args ...any) (db.Rows, error) {
	rows, err := t.tx.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	return &pgRows{rows: rows}, nil
}
func (t *pgTx) Exec(ctx context.Context, sql string, args ...any) error {
	_, err := t.tx.Exec(ctx, sql, args...)
	return err
}
func (t *pgTx) Begin(ctx context.Context) (db.Tx, error) {
	tx, err := t.tx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return &pgTx{tx: tx}, nil
}
func (t *pgTx) Commit(ctx context.Context) error   { return t.tx.Commit(ctx) }
func (t *pgTx) Rollback(ctx context.Context) error { return t.tx.Rollback(ctx) }
