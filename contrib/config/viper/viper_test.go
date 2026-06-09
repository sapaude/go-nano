package viperadapter_test

import (
	"os"
	"testing"

	viperadapter "github.com/sapaude/go-nano/contrib/config/viper"
)

func TestLoader(t *testing.T) {
	f, err := os.CreateTemp("", "config*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())
	f.WriteString("app:\n  name: test\n")
	f.Close()

	l, err := viperadapter.New(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	if v := l.Get("app.name"); v != "test" {
		t.Fatalf("expected test, got %v", v)
	}
}
