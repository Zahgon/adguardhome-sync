package client

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/bakito/adguardhome-sync/internal/client/model"
	"github.com/bakito/adguardhome-sync/internal/log"
	"github.com/bakito/adguardhome-sync/internal/types"
)

var l = log.GetLogger("client")

// New create a new api client.
func New(config types.AdGuardInstance) (Client, error) {
	_ = "STUB: not implemented"
	return *new(Client), nil
}

// #nosec G402 has to be explicitly enabled

func basicAuth(username, password string) string { _ = "STUB: not implemented"; return "" }

type apiClient struct {
	host   string
	client *model.AdguardHomeClient
	log    *zap.SugaredLogger
}

func (a apiClient) Host(context.Context) string { _ = "STUB: not implemented"; return "" }

func (a apiClient) GetServerStatus(ctx context.Context) (*model.ServerStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a apiClient) GetFilteringStatus(ctx context.Context) (*model.FilterStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a apiClient) SetFilteringConfig(ctx context.Context, config model.FilterConfig) error {
	_ = "STUB: not implemented"
	return nil
}

func write[B any](
	ctx context.Context,
	body B,
	req func(ctx context.Context, body B, reqEditors ...model.RequestEditorFn) (*http.Response, error),
) error {
	_ = "STUB: not implemented"
	return nil
}

func read[I any](
	ctx context.Context,
	req func(ctx context.Context, reqEditors ...model.RequestEditorFn) (*http.Response, error),
	parse func(rsp *http.Response) (*I, error),
) (*I, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func detailedError(resp *http.Response) error { _ = "STUB: not implemented"; return nil }
