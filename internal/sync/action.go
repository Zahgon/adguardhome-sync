package sync

import (
	"go.uber.org/zap"

	"github.com/bakito/adguardhome-sync/internal/client"
	"github.com/bakito/adguardhome-sync/internal/client/model"
	"github.com/bakito/adguardhome-sync/internal/types"
)

func setupActions(cfg *types.Config) (actions []syncAction) { _ = "STUB: not implemented"; return nil }

type syncAction interface {
	sync(ac *actionContext) error
	name() string
}

type actionContext struct {
	rl            *zap.SugaredLogger
	origin        *origin
	client        client.Client
	replicaStatus *model.ServerStatus
	replica       types.AdGuardInstance
	cfg           *types.Config
}

type defaultAction struct {
	myName string
	doSync func(ac *actionContext) error
}

func action(name string, f func(ac *actionContext) error) syncAction {
	_ = "STUB: not implemented"
	return *new(syncAction)
}

func (d *defaultAction) sync(ac *actionContext) error { _ = "STUB: not implemented"; return nil }

func (d *defaultAction) name() string { _ = "STUB: not implemented"; return "" }
