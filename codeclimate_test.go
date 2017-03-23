package codeclimate

import (
	"net/http"
	"testing"
)

func TestGetScore(t *testing.T) {
	c := NewClient("mykey", "myapp", &http.Client{})
	c.GetScore()
}
