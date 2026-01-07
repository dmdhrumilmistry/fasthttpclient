package client_test

import (
	"testing"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/valyala/fasthttp"
)

var rlclient = client.NewRateLimitedClient(100, 1, &fasthttp.Client{})

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
	Put(rlclient, t)
}

func TestRLCPatch(t *testing.T) {
	Patch(rlclient, t)
}

func TestRLCDelete(t *testing.T) {
	Delete(rlclient, t)
}
