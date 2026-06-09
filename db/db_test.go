package db_test

import (
	"context"
	"testing"

	"github.com/sapaude/go-nano/db"
)

// mockRows is a minimal Rows implementation for interface compliance tests.
type mockRows struct{}

func (m *mockRows) Next() bool          { return false }
func (m *mockRows) Scan(_ ...any) error { return nil }
func (m *mockRows) Close() error        { return nil }
func (m *mockRows) Err() error          { return nil }

// mockDB verifies interface compliance.
type mockDB struct{}

func (d *mockDB) Query(_ context.Context, _ string, _ ...any) (db.Rows, error) {
	return &mockRows{}, nil
}
func (d *mockDB) Exec(_ context.Context, _ string, _ ...any) error { return nil }
func (d *mockDB) Begin(_ context.Context) (db.Tx, error)           { return &mockTx{}, nil }

type mockTx struct{ mockDB }

func (t *mockTx) Commit(_ context.Context) error   { return nil }
func (t *mockTx) Rollback(_ context.Context) error { return nil }

func TestDBInterface(t *testing.T) {
	var _ db.DB = &mockDB{}
	var _ db.Tx = &mockTx{}
}
