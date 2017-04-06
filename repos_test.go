package codeclimate

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRepo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json := []byte(`{
      "data": {
        "attributes": {
          "score": 3.06,
          "analysis_version": 4,
          "branch": "master"
        }
      }
    }`)
		if r.Method != "GET" {
			// 1. Test for correct method
			t.Error(fmt.Sprintf("Expected method to be %s but got %s.", "GET", r.Method))
		}
		if r.RequestURI != "/v1/repos/myappid" {
			// 2. Test for correct path
			t.Error(fmt.Sprintf(
				"Expected path to be %s but got %s.",
				"/v1/repos/myappid/test_reports",
				r.RequestURI,
			))
		}
		if r.Header.Get("Authorization") != "Token token=myapikey" {
			// 3. Test presence of authorization header
			t.Error(fmt.Sprintf(
				"Expected Authorization header to be %s but got %s.",
				"Token token=myapikey",
				r.Header.Get("Authorization"),
			))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	}))
	defer ts.Close()
	c := NewClient("myapikey", "myappid", ts.URL, &http.Client{})
	repo, e := c.GetRepo()
	if e != nil {
		t.Error(e.Error())
	}

	// 4. Test that attributes are set correctly
	if repo.Body.Attributes.Score != 3.06 {
		t.Error(fmt.Sprintf("Expected Score to be %s but got %s", 3.06, repo.Body.Attributes.Score))
	}
	if repo.Body.Attributes.AnalysisVersion != 4 {
		t.Error(fmt.Sprintf("Expected AnalysisVersion to be %s but got %s", 4, repo.Body.Attributes.AnalysisVersion))
	}
	if repo.Body.Attributes.Branch != "master" {
		t.Error(fmt.Sprintf("Expected Branch to be %s but got %s", "master", repo.Body.Attributes.Branch))
	}
}
