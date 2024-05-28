package client_test

import (
	"testing"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
)

func TestGet(t *testing.T) {
	headers := map[string]string{
		"User-Agent": "fasthttpclient",
		"Accept":     "application/json",
	}

	queryParams := map[string]string{
		"accept": "json",
	}

	// testing function without providing query params and headers
	resp, err := client.Get("https://ipinfo.io", queryParams, headers)

	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

}
