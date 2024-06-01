package client_test

import (
	"testing"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
)

var fhclient = client.NewFHClient()

func TestFHCGet(t *testing.T) {
	Get(fhclient, t)
}
