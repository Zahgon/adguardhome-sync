package model

import (
	"go.uber.org/zap"
)

// Clone the config.
func (c *DhcpStatus) Clone() *DhcpStatus { _ = "STUB: not implemented"; return nil }

func (c *DhcpStatus) cleanV4V6() { _ = "STUB: not implemented"; return }

// CleanAndEquals dhcp server config equal check where V4 and V6 are cleaned in advance.
func (c *DhcpStatus) CleanAndEquals(o *DhcpStatus) bool { _ = "STUB: not implemented"; return false }

// Equals dhcp server config equal check.
func (c *DhcpStatus) Equals(o *DhcpStatus) bool { _ = "STUB: not implemented"; return false }

func (c *DhcpStatus) HasConfig() bool { _ = "STUB: not implemented"; return false }

func (j DhcpConfigV4) isValid() bool { _ = "STUB: not implemented"; return false }

func (j DhcpConfigV6) isValid() bool { _ = "STUB: not implemented"; return false }

type DhcpStaticLeases []DhcpStaticLease

// MergeDhcpStaticLeases the leases.
func MergeDhcpStaticLeases(l, other *[]DhcpStaticLease) (adds, removes DhcpStaticLeases) {
	_ = "STUB: not implemented"
	return *new(DhcpStaticLeases), *new(DhcpStaticLeases)
}

// Equals dns config equal check.
func (c *DNSConfig) Equals(o *DNSConfig) bool { _ = "STUB: not implemented"; return false }

func (c *DNSConfig) Clone() *DNSConfig { _ = "STUB: not implemented"; return nil }

// Sort dns config.
func (c *DNSConfig) Sort() { _ = "STUB: not implemented"; return }

// Equals access list equal check.
func (al *AccessList) Equals(o *AccessList) bool { _ = "STUB: not implemented"; return false }

func EqualsStringSlice(a, b *[]string, sortIt bool) bool { _ = "STUB: not implemented"; return false }

// Sort clients.
func (cl *Client) Sort() { _ = "STUB: not implemented"; return }

// PrepareDiff so we skip it in diff.
func (cl *Client) PrepareDiff() *string { _ = "STUB: not implemented"; return nil }

// AfterDiff reset after diff.
func (cl *Client) AfterDiff(tz *string) { _ = "STUB: not implemented"; return }

// Equals Clients equal check.
func (cl *Client) Equals(o *Client) bool { _ = "STUB: not implemented"; return false }

// Add ac client.
func (clients *Clients) Add(cl Client) { _ = "STUB: not implemented"; return }

// Merge merge Clients.
func (clients *Clients) Merge(other *Clients) (adds, removes, updates []*Client) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Key RewriteEntry key.
func (re *RewriteEntry) Key() string { _ = "STUB: not implemented"; return "" }

// RewriteEntries list of RewriteEntry.
type (
	RewriteEntries []RewriteEntry
	RewriteUpdates []RewriteUpdate
)

// Merge RewriteEntries.
func (rwe *RewriteEntries) Merge(other *RewriteEntries) (adds, removes, duplicates RewriteEntries, updates RewriteUpdates) {
	_ = "STUB: not implemented"
	return *new(RewriteEntries), *new(RewriteEntries), *new(RewriteEntries), *new(RewriteUpdates)
}

// remove duplicate

//	skip duplicate

func MergeFilters(this, other *[]Filter) (adds, updates, removes []Filter) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Equals Filter equal check.
func (f *Filter) Equals(o *Filter) bool { _ = "STUB: not implemented"; return false }

type QueryLogConfigWithIgnored struct {
	QueryLogConfig

	// Ignored List of host names, which should not be written to log
	Ignored []string `json:"ignored,omitempty"`
}

// Equals QueryLogConfig equal check.
func (qlc *QueryLogConfigWithIgnored) Equals(o *QueryLogConfigWithIgnored) bool {
	_ = "STUB: not implemented"
	return false
}

// Equals QueryLogConfigInterval equal check.
func (qlc *QueryLogConfigInterval) Equals(o *QueryLogConfigInterval) bool {
	_ = "STUB: not implemented"
	return false
}

func ptrEquals[T comparable](a, b *T) bool { _ = "STUB: not implemented"; return false }

// EnableConfig API struct.
type EnableConfig struct {
	Enabled bool `json:"enabled"`
}

func (ssc *SafeSearchConfig) Equals(o *SafeSearchConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func (pi *ProfileInfo) Equals(o *ProfileInfo, withTheme bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (pi *ProfileInfo) ShouldSyncFor(o *ProfileInfo, withTheme bool) *ProfileInfo {
	_ = "STUB: not implemented"
	return nil
}

func (bss *BlockedServicesSchedule) Equals(o *BlockedServicesSchedule) bool {
	_ = "STUB: not implemented"
	return false
}

func (bss *BlockedServicesSchedule) ServicesString() string { _ = "STUB: not implemented"; return "" }

func ArrayString(a *[]string) string { _ = "STUB: not implemented"; return "" }

func (c *DNSConfig) Sanitize(l *zap.SugaredLogger) {
	_ = "STUB: not implemented"
	// disable UsePrivatePtrResolvers if not configured
	// https://github.com/AdguardTeam/AdGuardHome/issues/6820
	return
}

// Equals GetStatsConfigResponse equal check.
func (sc *GetStatsConfigResponse) Equals(o *GetStatsConfigResponse) bool {
	_ = "STUB: not implemented"
	return false
}

func NewStats() *Stats { _ = "STUB: not implemented"; return nil }

func (s *Stats) Add(other *Stats) { _ = "STUB: not implemented"; return }

func addInt(t, add *int) *int { _ = "STUB: not implemented"; return nil }

func sumUp(t, o *[]int) *[]int { _ = "STUB: not implemented"; return nil }

func (c *TlsConfig) Equals(config *TlsConfig) bool { _ = "STUB: not implemented"; return false }
