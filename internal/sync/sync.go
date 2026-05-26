package sync

import (
	"time"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"

	"github.com/bakito/adguardhome-sync/internal/client"
	"github.com/bakito/adguardhome-sync/internal/client/model"
	"github.com/bakito/adguardhome-sync/internal/log"
	"github.com/bakito/adguardhome-sync/internal/types"
)

var l = log.GetLogger("sync")

// Sync config from origin to replica.
func Sync(cfg *types.Config) error { _ = "STUB: not implemented"; return nil }

func runOnStartAsync(cfg *types.Config, w *worker) { _ = "STUB: not implemented"; return }

type worker struct {
	cfg          *types.Config
	running      bool
	cron         *cron.Cron
	createClient func(instance types.AdGuardInstance, timeout time.Duration) (client.Client, error)
	actions      []syncAction
}

func (w *worker) status() *syncStatus { _ = "STUB: not implemented"; return nil }

func (w *worker) getStatus(inst types.AdGuardInstance) replicaStatus {
	_ = "STUB: not implemented"
	return *new(replicaStatus)
}

func (w *worker) sync() { _ = "STUB: not implemented"; return }

// Workaround for https://github.com/AdguardTeam/AdGuardHome/issues/7987
// and https://github.com/AdguardTeam/AdGuardHome/issues/7985

func (w *worker) syncTo(l *zap.SugaredLogger, o *origin, replica types.AdGuardInstance) {
	_ = "STUB: not implemented"
	return
}

func (*worker) statusWithSetup(
	rl *zap.SugaredLogger,
	replica types.AdGuardInstance,
	rc client.Client,
) (*model.ServerStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type origin struct {
	status                  *model.ServerStatus
	rewrites                *model.RewriteEntries
	blockedServicesSchedule *model.BlockedServicesSchedule
	filters                 *model.FilterStatus
	clients                 *model.Clients
	queryLogConfig          *model.QueryLogConfigWithIgnored
	statsConfig             *model.GetStatsConfigResponse
	accessList              *model.AccessList
	dnsConfig               *model.DNSConfig
	dhcpServerConfig        *model.DhcpStatus
	parental                bool
	safeSearch              *model.SafeSearchConfig
	profileInfo             *model.ProfileInfo
	safeBrowsing            bool
	tlsConfig               *model.TlsConfig
}
