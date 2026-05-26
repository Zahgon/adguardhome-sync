package types

import (
	"go.uber.org/zap"
)

func NewFeatures(enabled bool) Features { _ = "STUB: not implemented"; return *new(Features) }

// Features feature flags.
type Features struct {
	DNS              DNS         `json:"dns"              yaml:"dns"`
	DHCP             DHCP        `json:"dhcp"             yaml:"dhcp"`
	GeneralSettings  bool        `json:"generalSettings"  yaml:"generalSettings"  documentation:"Sync general settings"                                                env:"FEATURES_GENERAL_SETTINGS"`
	ProtectionStatus bool        `json:"protectionStatus" yaml:"protectionStatus" documentation:"Sync the protection status (disabled if generalSettings is disabled)" env:"FEATURES_PROTECTION_STATUS"`
	QueryLogConfig   bool        `json:"queryLogConfig"   yaml:"queryLogConfig"   documentation:"Sync query log config"                                                env:"FEATURES_QUERY_LOG_CONFIG"`
	StatsConfig      bool        `json:"statsConfig"      yaml:"statsConfig"      documentation:"Sync stats config"                                                    env:"FEATURES_STATS_CONFIG"`
	ClientSettings   bool        `json:"clientSettings"   yaml:"clientSettings"   documentation:"Sync client settings"                                                 env:"FEATURES_CLIENT_SETTINGS"`
	Services         bool        `json:"services"         yaml:"services"         documentation:"Sync services"                                                        env:"FEATURES_SERVICES"`
	Filters          FiltersType `json:"filters"          yaml:"filters"          documentation:"Sync filters (use sub-fields for granular control)"                   env:"FEATURES_FILTERS"`
	Theme            bool        `json:"theme"            yaml:"theme"            documentation:"Sync the web UI theme"                                                env:"FEATURES_THEME"`
	TLSConfig        bool        `json:"tlsConfig"        yaml:"tlsConfig"        documentation:"Sync the TLS config"                                                  env:"FEATURES_TLS_CONFIG"`
}

// FiltersType features.
type FiltersType struct {
	Blacklist bool `documentation:"Sync blacklist filters" env:"FEATURES_FILTERS_BLACKLIST"    json:"blacklist" yaml:"blacklist"`
	Whitelist bool `documentation:"Sync whitelist filters" env:"FEATURES_FILTERS_WHITELIST"    json:"whitelist" yaml:"whitelist"`
	UserRules bool `documentation:"Sync user rules"        env:"FEATURES_FILTERS_USER_RULES"   json:"userRules" yaml:"userRules"`
}

// UnmarshalYAML implements custom unmarshalling for FiltersType.
func (f *FiltersType) UnmarshalYAML(unmarshal func(any) error) error {
	_ = "STUB: not implemented"
	return nil
}

// UnmarshalJSON implements custom unmarshalling for FiltersType.
func (f *FiltersType) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalText implements custom unmarshalling for env vars.
func (f *FiltersType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

// DHCP features.
type DHCP struct {
	ServerConfig bool `documentation:"Sync DHCP server config" env:"FEATURES_DHCP_SERVER_CONFIG" json:"serverConfig" yaml:"serverConfig"`
	StaticLeases bool `documentation:"Sync DHCP static leases" env:"FEATURES_DHCP_STATIC_LEASES" json:"staticLeases" yaml:"staticLeases"`
}

// DNS features.
type DNS struct {
	AccessLists  bool `documentation:"Sync DNS access lists"  env:"FEATURES_DNS_ACCESS_LISTS"  json:"accessLists"  yaml:"accessLists"`
	ServerConfig bool `documentation:"Sync DNS server config" env:"FEATURES_DNS_SERVER_CONFIG" json:"serverConfig" yaml:"serverConfig"`
	Rewrites     bool `documentation:"Sync DNS rewrites"      env:"FEATURES_DNS_REWRITES"      json:"rewrites"     yaml:"rewrites"`
}

// LogDisabled log all disabled features.
func (f *Features) LogDisabled(l *zap.SugaredLogger) { _ = "STUB: not implemented"; return }

func (f *Features) collectDisabled() []string { _ = "STUB: not implemented"; return nil }
