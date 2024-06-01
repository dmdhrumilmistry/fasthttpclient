package client_test

import (
	"testing"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
)

var rlclient = client.NewRateLimitedClient(100, 1)

func TestRLCGet(t *testing.T) {
	Get(rlclient, t)
}

func TestRLCHead(t *testing.T) {
	Head(rlclient, t)
}

func TestRLCPost(t *testing.T) {
	Post(rlclient, t)
}

func TestRLCPut(t *testing.T) {
	Put(fhclient, t)
}

func TestRLCPatch(t *testing.T) {
	Patch(fhclient, t)
}

func TestRLCDelete(t *testing.T) {
	Delete(fhclient, t)
}
