package client

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
)

func (cl *client) doGet(req *resty.Request, url string) error {
	_ = "STUB: not implemented"
	return nil
}

func (cl *client) doPost(req *resty.Request, url string) error {
	_ = "STUB: not implemented"
	return nil
}

// adguard home requires content type json to be set on every request

func (cl *client) doPut(req *resty.Request, url string) error {
	_ = "STUB: not implemented"
	return nil
}

// adguard home requires content type json to be set on every request

func checkAuthenticationIssue(resp *resty.Response, rl *zap.SugaredLogger) {
	_ = "STUB: not implemented"
	return
}
