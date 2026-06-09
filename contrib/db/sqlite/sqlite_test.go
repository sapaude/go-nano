package sqliteadapter_test

import (
	"context"
	"testing"

	sqliteadapter "github.com/sapaude/go-nano/contrib/db/sqlite"
)

func TestSQLite(t *testing.T) {
	d := sqliteadapter.New(sqliteadapter.Config{Path: ":memory:"})
	if d.Name() != "db.sqlite" {
		t.Fatalf("unexpected name: %s", d.Name())
	}
	ctx := context.Background()
	if err := d.Init(ctx); err != nil {
		t.Fatal(err)
	}
	defer d.Stop(ctx)

	if err := d.Exec(ctx, "CREATE TABLE t (id INTEGER PRIMARY KEY)"); err != nil {
		t.Fatal(err)
	}
	rows, err := d.Query(ctx, "SELECT count(*) FROM t")
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
}
