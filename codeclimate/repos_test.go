package codeclimate

import (
	"net/http"
	"testing"
)

func TestGetRepo(t *testing.T) {
	c := NewClient(
		"ee89d8df16816e4243fd145b358793f0321d845d",
		"53ab23a7e30ba070f000ac26",
		&http.Client{},
	)
	c.GetRepo()
}
