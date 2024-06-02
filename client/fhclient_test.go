package client_test

import (
	"testing"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/valyala/fasthttp"
)

var fhclient = client.NewFHClient(&fasthttp.Client{})

func TestFHCGet(t *testing.T) {
	Get(fhclient, t)
}

func TestFHCHead(t *testing.T) {
	Head(fhclient, t)
}

func TestFHCPost(t *testing.T) {
	Post(fhclient, t)
}

func TestFHCPut(t *testing.T) {
	Put(fhclient, t)
}

func TestFHCPatch(t *testing.T) {
	Patch(fhclient, t)
}

func TestFHCDelete(t *testing.T) {
	Delete(fhclient, t)
}
