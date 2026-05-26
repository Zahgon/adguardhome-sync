package sync

import (
	"github.com/gin-gonic/gin"

	// go embed blank import.
	_ "embed"
)

func (w *worker) handleSync(c *gin.Context) { _ = "STUB: not implemented"; return }

func (w *worker) handleRoot(c *gin.Context) { _ = "STUB: not implemented"; return }

func percent(a, b *int) string { _ = "STUB: not implemented"; return "" }

func (*worker) handleLogs(c *gin.Context) { _ = "STUB: not implemented"; return }

func (*worker) handleClearLogs(c *gin.Context) { _ = "STUB: not implemented"; return }

func (w *worker) handleStatus(c *gin.Context) { _ = "STUB: not implemented"; return }

func (w *worker) handleHealthz(c *gin.Context) { _ = "STUB: not implemented"; return }

func (w *worker) listenAndServe() { _ = "STUB: not implemented"; return }

// kill -SIGHUP XXXX
// kill -SIGINT XXXX or Ctrl+c
// kill -SIGQUIT XXXX

// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)

type syncStatus struct {
	SyncRunning bool            `json:"syncRunning"`
	Origin      replicaStatus   `json:"origin"`
	Replicas    []replicaStatus `json:"replicas"`
}

type replicaStatus struct {
	Host              string `json:"host"`
	URL               string `json:"url"`
	Status            string `json:"status"`
	Error             string `json:"error,omitempty"`
	ProtectionEnabled *bool  `json:"protection_enabled"`
}

func getLast24Hours() []string { _ = "STUB: not implemented"; return nil }

// Loop to get the last 24 hours

// Calculate the time for the current hour in the loop

// Format the time as "14 Dec 17:00"

// Reverse the slice to get the correct order (from oldest to latest)
