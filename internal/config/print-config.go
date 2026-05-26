package config

import (
	"github.com/bakito/adguardhome-sync/internal/types"

	_ "embed"
)

//go:embed print-config.md
var printConfigTemplate string

func (ac *AppConfig) Print() error { _ = "STUB: not implemented"; return nil }

func aghVersion(i types.AdGuardInstance) string { _ = "STUB: not implemented"; return "" }

func (ac *AppConfig) printInternal(env []string, originVersion string, replicaVersions []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// The name "inc" is what the function will be called in the template text.
