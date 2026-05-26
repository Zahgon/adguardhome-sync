package config

import (
	"regexp"

	"github.com/bakito/adguardhome-sync/internal/log"
	"github.com/bakito/adguardhome-sync/internal/types"
)

var (
	envReplicasURLPattern = regexp.MustCompile(`^REPLICA(\d+)_URL=(.*)`)
	logger                = log.GetLogger("config")
)

type AppConfig struct {
	cfg      *types.Config
	filePath string
	content  string
}

func (ac *AppConfig) PrintConfigOnly() bool { _ = "STUB: not implemented"; return false }

func (ac *AppConfig) Get() *types.Config { _ = "STUB: not implemented"; return nil }

func (ac *AppConfig) Init() error { _ = "STUB: not implemented"; return nil }

func Get(configFile string, flags Flags) (*AppConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// read yaml config

// overwrite from command flags

// *bool field creates issues when already not nil
// origin filed makes no sense to be set.

// keep previously set value

// ignore origin and replicas form env parsing as they are handled separately

// overwrite from env vars

// restore origin and replica

// if not set from env, use previous value

func initialConfig() *types.Config { _ = "STUB: not implemented"; return nil }
