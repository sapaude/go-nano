package viperadapter

import (
	"fmt"

	"github.com/spf13/viper"
)

// Loader loads configuration from a file using Viper.
type Loader struct {
	v *viper.Viper
}

// New creates a new Loader with the given config file path.
func New(configFile string) (*Loader, error) {
	v := viper.New()
	v.SetConfigFile(configFile)
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	return &Loader{v: v}, nil
}

// Unmarshal unmarshals the given key into dest.
func (l *Loader) Unmarshal(key string, dest any) error {
	return l.v.UnmarshalKey(key, dest)
}

// Get returns the value for the given key.
func (l *Loader) Get(key string) any {
	return l.v.Get(key)
}
