package config

import (
	"github.com/bakito/adguardhome-sync/internal/types"
)

// Manually collect replicas from env.
func enrichReplicasFromEnv(initialReplicas []types.AdGuardInstance) ([]types.AdGuardInstance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// keep the previously set value
