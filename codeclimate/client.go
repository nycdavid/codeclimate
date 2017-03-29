package codeclimate

import (
	"net/http"
)

const (
	codeClimateApiHost = "https://api.codeclimate.com"
)

type HttpCaller interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	apiKey     string
	appId      string
	httpClient HttpCaller
}

func NewClient(apiKey string, appId string, caller HttpCaller) client {
	return client{
		apiKey:     apiKey,
		appId:      appId,
		httpClient: caller,
	}
}
