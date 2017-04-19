package codeclimate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type TestReports struct {
	Body []TestReport `json:"data"`
}

type TestReport struct {
	Id         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		CoveredPercent float64 `json:"covered_percent"`
	} `json:"attributes"`
}

func (c client) GetTestReport() (TestReport, error) {
	var trs TestReports
	u, e := url.Parse(c.BaseUrl)
	if e != nil {
		return TestReport{}, e
	}
	u.Path = fmt.Sprintf("/repos/%s/test_reports", c.AppId)
	req, e := http.NewRequest(http.MethodGet, u.String(), nil)
	if e != nil {
		return TestReport{}, e
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", c.ApiKey))
	httpres, e := c.httpClient.Do(req)
	if e != nil {
		return TestReport{}, e
	}
	dec := json.NewDecoder(httpres.Body)
	e = dec.Decode(&trs)
	if e != nil {
		fmt.Println(e.Error())
	}
	return trs.Body[0], nil
}
