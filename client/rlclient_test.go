package client_test

import (
	"testing"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
)

var rlclient = client.NewRateLimitedClient(100, 1)

func TestRLCGet(t *testing.T) {
	Get(rlclient, t)
}
