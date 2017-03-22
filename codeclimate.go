package codeclimate

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	codeClimateApiHost = "https://api.codeclimate.com"
)

type HttpCaller interface {
	Do(req *http.Request) (*http.Response, error)
}

type codeClimateClient struct {
	apiKey     string
	appId      string
	httpClient HttpCaller
}

func New(apiKey string, appId string) codeClimateClient {
	return codeClimateClient{
		apiKey:     apiKey,
		appId:      appId,
		httpClient: http.Client{},
	}
}

func (c codeClimateClient) GetScore() float64 {
	u = url.Parse(codeClimateApiHost)
	u.Path = fmt.Sprintf("/v1/repos/%s", c.appId)
}
