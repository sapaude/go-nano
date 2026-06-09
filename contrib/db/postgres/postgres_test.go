package pgadapter_test

import (
	"testing"

	pgadapter "github.com/sapaude/go-nano/contrib/db/postgres"
)

func TestNew(t *testing.T) {
	d := pgadapter.New(pgadapter.Config{DSN: "postgres://localhost/test", MaxConns: 10})
	if d.Name() != "db.postgres" {
		t.Fatalf("unexpected name: %s", d.Name())
	}
}
