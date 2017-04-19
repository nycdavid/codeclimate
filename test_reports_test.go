package codeclimate

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTestReport(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			// 1. Test for correct method
			t.Error(fmt.Sprintf("Expected method to be %s but got %s.", "GET", r.Method))
		}
		if r.RequestURI != "/repos/myappid/test_reports" {
			// 2. Test for correct path
			t.Error(fmt.Sprintf(
				"Expected path to be %s but got %s.",
				"/repos/myappid/test_reports",
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
		w.Write([]byte(`
      {
        "data": [{
          "id": "faketestreportid",
          "type": "test_report",
          "attributes": {
            "covered_percent": 97.03
          }
        }]
      }
    `))
	}))
	defer ts.Close()
	c := NewClient("myapikey", "myappid", ts.URL, &http.Client{})
	tr, e := c.GetTestReport()
	if e != nil {
		t.Error(e.Error())
	}

	// 4. Test that attributes are set correctly
	if tr.Id != "faketestreportid" {
		t.Error(fmt.Sprintf("Expected tr.Id to be %s but got %s", "faketestreportid", tr.Id))
	}
	if tr.Type != "test_report" {
		t.Error(fmt.Sprintf("Expected tr.Type to be %s but got %s", "test_report", tr.Type))
	}
	if tr.Attributes.CoveredPercent != 97.03 {
		t.Error(fmt.Sprintf("Expected tr.Attributes.CoveredPercent to be %s but got %s", "97.03", tr.Attributes.CoveredPercent))
	}
}
