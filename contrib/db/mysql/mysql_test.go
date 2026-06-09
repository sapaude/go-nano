package mysqladapter_test

import (
	"testing"

	mysqladapter "github.com/sapaude/go-nano/contrib/db/mysql"
)

func TestNew(t *testing.T) {
	d := mysqladapter.New(mysqladapter.Config{DSN: "user:pass@tcp(localhost:3306)/test"})
	if d.Name() != "db.mysql" {
		t.Fatalf("unexpected name: %s", d.Name())
	}
}
