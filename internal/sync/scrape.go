package sync

import (
	"github.com/bakito/adguardhome-sync/internal/metrics"
	"github.com/bakito/adguardhome-sync/internal/types"
)

func (w *worker) startScraping() { _ = "STUB: not implemented"; return }

func (w *worker) scrape() { _ = "STUB: not implemented"; return }

func (w *worker) getMetrics(inst types.AdGuardInstance) metrics.InstanceMetrics {
	_ = "STUB: not implemented"
	return *new(metrics.InstanceMetrics)
}
