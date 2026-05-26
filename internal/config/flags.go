package config

import (
	"github.com/bakito/adguardhome-sync/internal/types"
)

func readFlags(cfg *types.Config, flags Flags) error { _ = "STUB: not implemented"; return nil }

type flagReader struct {
	cfg   *types.Config
	flags Flags
}

func (fr *flagReader) readReplicaFlags() error { _ = "STUB: not implemented"; return nil }

func (fr *flagReader) readOriginFlags() error { _ = "STUB: not implemented"; return nil }

func (fr *flagReader) readFeatureFlags() error { _ = "STUB: not implemented"; return nil }

func (fr *flagReader) readAPIFlags() error { _ = "STUB: not implemented"; return nil }

func (fr *flagReader) readRootFlags() error { _ = "STUB: not implemented"; return nil }

type Flags interface {
	Changed(name string) bool
	GetString(name string) (string, error)
	GetInt(name string) (int, error)
	GetBool(name string) (bool, error)
}

func (fr *flagReader) setStringFlag(name string, cb callback[string]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (fr *flagReader) setBoolFlag(name string, cb callback[bool]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (fr *flagReader) setIntFlag(name string, cb callback[int]) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type callback[T any] func(_ *types.Config, value T)
