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
			t.Error(fmt.Sprintf("Expected method to be %s but got %s.", "GET", r.Method))
		}
		// 2. Test path
		// 3. Test presence of authorization header
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
