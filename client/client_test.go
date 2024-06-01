package client_test

import (
	"testing"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
)

// test constants
const jsonContentTypeHeader = "application/json"

const baseEndpoint string = "https://jsonplaceholder.typicode.com"
const postEndpoint string = baseEndpoint + "/posts"
const postsResourceEndpoint string = postEndpoint + "/1"

// variables for testing
var body = []byte(`{"title":"fasthttpclient api test","body":"this is a test message","userId":1}`)
var queryParams = map[string]string{"accept": "json"}
var headers = map[string]string{
	"User-Agent":   "fasthttpclient",
	"Accept":       jsonContentTypeHeader,
	"Content-Type": jsonContentTypeHeader,
}

func Get(testClient client.ClientInterface, t *testing.T) {
	resp, err := client.Get(testClient, postsResourceEndpoint, queryParams, headers)

	if err != nil {
		t.Error(err)
	}

	StatusCodeTest(resp, 200, t)
}

func Head(testClient client.ClientInterface, t *testing.T) {
	resp, err := client.Head(testClient, postsResourceEndpoint, queryParams, headers)

	if err != nil {
		t.Error(err)
	}

	StatusCodeTest(resp, 200, t)
}

func Post(testClient client.ClientInterface, t *testing.T) {
	resp, err := client.Post(testClient, postEndpoint, queryParams, headers, body)

	if err != nil {
		t.Error(err)
	}

	StatusCodeTest(resp, 201, t)
}

func Put(testClient client.ClientInterface, t *testing.T) {
	resp, err := client.Put(testClient, postsResourceEndpoint, queryParams, headers, body)

	if err != nil {
		t.Error(err)
	}

	StatusCodeTest(resp, 200, t)
}

func Patch(testClient client.ClientInterface, t *testing.T) {
	resp, err := client.Patch(testClient, postsResourceEndpoint, queryParams, headers, body)

	if err != nil {
		t.Error(err)
	}

	StatusCodeTest(resp, 200, t)
}

func Delete(testClient client.ClientInterface, t *testing.T) {
	resp, err := client.Delete(testClient, postsResourceEndpoint, queryParams, headers, nil)

	if err != nil {
		t.Error(err)
	}

	StatusCodeTest(resp, 200, t)
}
