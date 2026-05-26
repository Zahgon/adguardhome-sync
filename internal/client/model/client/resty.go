package client

import (
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/bakito/adguardhome-sync/internal/client/model"
)

var _ model.HttpRequestDoer = &adapter{}

func RestyAdapter(r *resty.Client) model.HttpRequestDoer {
	_ = "STUB: not implemented"
	return *new(model.HttpRequestDoer)
}

type adapter struct {
	client *resty.Client
}

func (a adapter) Do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
