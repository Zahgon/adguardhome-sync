package client

import (
	"errors"
	"time"

	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"

	"github.com/bakito/adguardhome-sync/internal/client/model"
	"github.com/bakito/adguardhome-sync/internal/log"
	"github.com/bakito/adguardhome-sync/internal/types"
)

const envRedirectPolicyNoOfRedirects = "REDIRECT_POLICY_NO_OF_REDIRECTS"

type Error struct {
	message   string
	errorCode int
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Error) Code() int { _ = "STUB: not implemented"; return 0 }

var (
	l = log.GetLogger("client")
	// ErrSetupNeeded custom error.
	ErrSetupNeeded = errors.New("setup needed")
)

func detailedError(resp *resty.Response, err error) error { _ = "STUB: not implemented"; return nil }

// New create a new client.
func New(config types.AdGuardInstance, timeout time.Duration) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// #nosec G402 has to be explicitly enabled

// no redirect

// Client AdguardHome API client interface.
//
//nolint:interfacebloat
type Client interface {
	Host() string
	Status() (*model.ServerStatus, error)
	Stats() (*model.Stats, error)
	QueryLog(limit int) (*model.QueryLog, error)
	ToggleProtection(enable bool) error
	RewriteList() (*model.RewriteEntries, error)
	AddRewriteEntries(e ...model.RewriteEntry) error
	DeleteRewriteEntries(e ...model.RewriteEntry) error
	UpdateRewriteEntries(e ...model.RewriteUpdate) error
	Filtering() (*model.FilterStatus, error)
	ToggleFiltering(enabled bool, interval int) error
	AddFilter(whitelist bool, f model.Filter) error
	DeleteFilter(whitelist bool, f model.Filter) error
	UpdateFilter(whitelist bool, f model.Filter) error
	RefreshFilters(whitelist bool) error
	SetCustomRules(rules *[]string) error
	SafeBrowsing() (bool, error)
	ToggleSafeBrowsing(enable bool) error
	Parental() (bool, error)
	ToggleParental(enable bool) error
	SafeSearchConfig() (*model.SafeSearchConfig, error)
	SetSafeSearchConfig(settings *model.SafeSearchConfig) error
	ProfileInfo() (*model.ProfileInfo, error)
	SetProfileInfo(settings *model.ProfileInfo) error
	BlockedServicesSchedule() (*model.BlockedServicesSchedule, error)
	SetBlockedServicesSchedule(schedule *model.BlockedServicesSchedule) error
	Clients() (*model.Clients, error)
	AddClient(client *model.Client) error
	UpdateClient(client *model.Client) error
	DeleteClient(client *model.Client) error
	QueryLogConfig() (*model.QueryLogConfigWithIgnored, error)
	SetQueryLogConfig(ql *model.QueryLogConfigWithIgnored) error
	StatsConfig() (*model.GetStatsConfigResponse, error)
	SetStatsConfig(sc *model.PutStatsConfigUpdateRequest) error
	Setup() error
	AccessList() (*model.AccessList, error)
	SetAccessList(accessList *model.AccessList) error
	DNSConfig() (*model.DNSConfig, error)
	SetDNSConfig(config *model.DNSConfig) error
	DhcpConfig() (*model.DhcpStatus, error)
	SetDhcpConfig(status *model.DhcpStatus) error
	AddDHCPStaticLease(lease model.DhcpStaticLease) error
	DeleteDHCPStaticLease(lease model.DhcpStaticLease) error
	TLSConfig() (*model.TlsConfig, error)
	SetTLSConfig(tls *model.TlsConfig) error
}

type client struct {
	client  *resty.Client
	log     *zap.SugaredLogger
	host    string
	version string
}

func (cl *client) Host() string { _ = "STUB: not implemented"; return "" }

func contentType(resp *resty.Response) string { _ = "STUB: not implemented"; return "" }

func (cl *client) Status() (*model.ServerStatus, error) { _ = "STUB: not implemented"; return nil, nil }

func (cl *client) Stats() (*model.Stats, error) { _ = "STUB: not implemented"; return nil, nil }

func (cl *client) QueryLog(limit int) (*model.QueryLog, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *client) RewriteList() (*model.RewriteEntries, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *client) AddRewriteEntries(entries ...model.RewriteEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) DeleteRewriteEntries(entries ...model.RewriteEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) UpdateRewriteEntries(entries ...model.RewriteUpdate) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) SafeBrowsing() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (cl *client) ToggleSafeBrowsing(enable bool) error { _ = "STUB: not implemented"; return nil }

func (cl *client) Parental() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (cl *client) ToggleParental(enable bool) error { _ = "STUB: not implemented"; return nil }

func (cl *client) toggleStatus(mode string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (cl *client) toggleBool(mode string, enable bool) error { _ = "STUB: not implemented"; return nil }

func (cl *client) Filtering() (*model.FilterStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *client) AddFilter(whitelist bool, f model.Filter) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) DeleteFilter(whitelist bool, f model.Filter) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) UpdateFilter(whitelist bool, f model.Filter) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) RefreshFilters(whitelist bool) error { _ = "STUB: not implemented"; return nil }

func (cl *client) ToggleProtection(enable bool) error { _ = "STUB: not implemented"; return nil }

func (cl *client) SetCustomRules(rules *[]string) error { _ = "STUB: not implemented"; return nil }

func (cl *client) ToggleFiltering(enabled bool, interval int) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) BlockedServicesSchedule() (*model.BlockedServicesSchedule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *client) SetBlockedServicesSchedule(schedule *model.BlockedServicesSchedule) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) Clients() (*model.Clients, error) { _ = "STUB: not implemented"; return nil, nil }

func (cl *client) AddClient(client *model.Client) error { _ = "STUB: not implemented"; return nil }

func (cl *client) UpdateClient(client *model.Client) error { _ = "STUB: not implemented"; return nil }

func (cl *client) DeleteClient(client *model.Client) error { _ = "STUB: not implemented"; return nil }

func (cl *client) QueryLogConfig() (*model.QueryLogConfigWithIgnored, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *client) SetQueryLogConfig(qlc *model.QueryLogConfigWithIgnored) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) StatsConfig() (*model.GetStatsConfigResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *client) SetStatsConfig(sc *model.PutStatsConfigUpdateRequest) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) Setup() error { _ = "STUB: not implemented"; return nil }

func (cl *client) AccessList() (*model.AccessList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *client) SetAccessList(list *model.AccessList) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) DNSConfig() (*model.DNSConfig, error) { _ = "STUB: not implemented"; return nil, nil }

func (cl *client) SetDNSConfig(config *model.DNSConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) DhcpConfig() (*model.DhcpStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *client) SetDhcpConfig(config *model.DhcpStatus) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) AddDHCPStaticLease(l model.DhcpStaticLease) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) DeleteDHCPStaticLease(l model.DhcpStaticLease) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) SafeSearchConfig() (*model.SafeSearchConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *client) SetSafeSearchConfig(settings *model.SafeSearchConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) ProfileInfo() (*model.ProfileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cl *client) SetProfileInfo(profile *model.ProfileInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) TLSConfig() (*model.TlsConfig, error) { _ = "STUB: not implemented"; return nil, nil }

func (cl *client) SetTLSConfig(tlsc *model.TlsConfig) error { _ = "STUB: not implemented"; return nil }
