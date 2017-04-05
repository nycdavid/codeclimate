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
	ApiKey     string
	AppId      string
	httpClient HttpCaller
	BaseUrl    string
}

func NewClient(apiKey string, appId string, baseUrl string, caller HttpCaller) client {
	return client{
		ApiKey:     apiKey,
		AppId:      appId,
		httpClient: caller,
		BaseUrl:    baseUrl,
	}
}
