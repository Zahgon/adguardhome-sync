package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/bakito/adguardhome-sync/internal/client/model"
	"github.com/bakito/adguardhome-sync/internal/log"
)

const StatsTotal = "total"

var (
	l = log.GetLogger("metrics")

	// avgProcessingTime - Average processing time for a DNS query.
	avgProcessingTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "avg_processing_time",
			Namespace: "adguard",
			Help:      "This represent the average processing time for a DNS query in s",
		},
		[]string{"hostname"},
	)

	// dnsQueries - Number of DNS queries.
	dnsQueries = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "num_dns_queries",
			Namespace: "adguard",
			Help:      "Number of DNS queries",
		},
		[]string{"hostname"},
	)

	// blockedFiltering - Number of DNS queries blocked.
	blockedFiltering = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "num_blocked_filtering",
			Namespace: "adguard",
			Help:      "This represent the number of domains blocked",
		},
		[]string{"hostname"},
	)

	// parentalFiltering - Number of DNS queries replaced by parental control.
	parentalFiltering = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "num_replaced_parental",
			Namespace: "adguard",
			Help:      "This represent the number of domains blocked (parental)",
		},
		[]string{"hostname"},
	)

	// safeBrowsingFiltering - Number of DNS queries replaced by safe browsing.
	safeBrowsingFiltering = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "num_replaced_safebrowsing",
			Namespace: "adguard",
			Help:      "This represent the number of domains blocked (safe browsing)",
		},
		[]string{"hostname"},
	)

	// safeSearchFiltering - Number of DNS queries replaced by safe search.
	safeSearchFiltering = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "num_replaced_safesearch",
			Namespace: "adguard",
			Help:      "This represent the number of domains blocked (safe search)",
		},
		[]string{"hostname"},
	)

	// topQueries - The number of top queries.
	topQueries = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "top_queried_domains",
			Namespace: "adguard",
			Help:      "This represent the top queried domains",
		},
		[]string{"hostname", "domain"},
	)

	// topBlocked - The number of top domains blocked.
	topBlocked = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "top_blocked_domains",
			Namespace: "adguard",
			Help:      "This represent the top bloacked domains",
		},
		[]string{"hostname", "domain"},
	)

	// topClients - The number of top clients.
	topClients = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "top_clients",
			Namespace: "adguard",
			Help:      "This represent the top clients",
		},
		[]string{"hostname", "client"},
	)

	// queryTypes - The type of DNS Queries (A, AAAA...)
	queryTypes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "query_types",
			Namespace: "adguard",
			Help:      "This represent the DNS query types",
		},
		[]string{"hostname", "type"},
	)

	// running - If Adguard is running.
	running = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "running",
			Namespace: "adguard",
			Help:      "This represent if Adguard is running",
		},
		[]string{"hostname"},
	)

	// protectionEnabled - If Adguard protection is enabled.
	protectionEnabled = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "protection_enabled",
			Namespace: "adguard",
			Help:      "This represent if Adguard Protection is enabled",
		},
		[]string{"hostname"},
	)
	// aghsSyncDuration - the sync curation in seconds.
	aghsSyncDuration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "sync_duration_seconds",
			Namespace: "adguard_home_sync",
			Help:      "This represents the duration of the last sync in seconds",
		},
		[]string{"hostname"},
	)
	// aghsSyncSuccessful - the sync result.
	aghsSyncSuccessful = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "sync_successful",
			Namespace: "adguard_home_sync",
			Help:      "This represents the whether the last sync was successful",
		},
		[]string{"hostname"},
	)
	stats = OverallStats{}
)

// Init initializes all Prometheus metrics made available by AdGuard  exporter.
func Init() { _ = "STUB: not implemented"; return }

func initMetric(name string, metric *prometheus.GaugeVec) { _ = "STUB: not implemented"; return }

func UpdateInstances(iml InstanceMetricsList) { _ = "STUB: not implemented"; return }

func UpdateResult(host string, ok bool, duration float64) { _ = "STUB: not implemented"; return }

func updateMetrics(im InstanceMetrics) {
	_ = "STUB: not implemented"
	// Status
	return
}

// Stats

// LogQuery

type InstanceMetricsList struct {
	Metrics []InstanceMetrics `faker:"slice_len=5"`
}

type InstanceMetrics struct {
	HostName string
	Status   *model.ServerStatus
	Stats    *model.Stats
	QueryLog *model.QueryLog
}

type OverallStats map[string]*model.Stats

func (os OverallStats) consolidate() OverallStats {
	_ = "STUB: not implemented"
	return *new(OverallStats)
}

func safeMetric[T int | float64 | float32](v *T) float64 { _ = "STUB: not implemented"; return 0 }

func getStats() OverallStats { _ = "STUB: not implemented"; return *new(OverallStats) }

func (os OverallStats) Total() *model.Stats { _ = "STUB: not implemented"; return nil }
