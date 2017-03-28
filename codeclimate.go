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

type client struct {
	apiKey     string
	appId      string
	httpClient HttpCaller
}

type Response struct {
	Data struct {
		Repo Repo
	}
}

type Repo struct {
	Score float64
}

func NewClient(apiKey string, appId string, caller HttpCaller) client {
	return client{
		apiKey:     apiKey,
		appId:      appId,
		httpClient: caller,
	}
}

func (c client) GetRepo() (Response, error) {
	var app Response
	u, e := url.Parse(codeClimateApiHost)
	if e != nil {
		return Response{}, e
	}
	u.Path = fmt.Sprintf("/v1/repos/%s", c.appId)
	req, e := http.NewRequest(http.MethodGet, u.String(), nil)
	if e != nil {
		return Response{}, e
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", c.apiKey))
	res, e := c.httpClient.Do(req)
	if e != nil {
		return Response{}, e
	}
	dec := json.NewDecoder(res.Body)
	e = dec.Decode(&app)
	fmt.Println(app.Data)
	if e != nil {
		fmt.Println(e.Error())
	}
	return app, nil
}
