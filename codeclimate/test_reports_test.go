package codeclimate

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTestReport(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	defer ts.Close()
	c := NewClient("myapikey", "myappid", ts.URL, &http.Client{})
}
