package codeclimate

import (
	"encoding/json"
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

func NewClient(apiKey string, appId string, caller HttpCaller) codeClimateClient {
	return codeClimateClient{
		apiKey:     apiKey,
		appId:      appId,
		httpClient: caller,
	}
}

type App struct {
	Data map[string]interface{} `json:"data"`
}

func (c codeClimateClient) GetScore() (float64, error) {
	var app App
	u, e := url.Parse(codeClimateApiHost)
	if e != nil {
		return 0.0, e
	}
	u.Path = fmt.Sprintf("/v1/repos/%s", c.appId)
	req, e := http.NewRequest(http.MethodGet, u.String(), nil)
	if e != nil {
		return 0.0, e
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", c.apiKey))
	res, e := c.httpClient.Do(req)
	if e != nil {
		return 0.0, e
	}
	dec := json.NewDecoder(res.Body)
	e = dec.Decode(&app)
	if e != nil {
		return 0.0, e
	}
	data := app.Data["attributes"].(map[string]interface{})
	return data["score"].(float64), nil
}
