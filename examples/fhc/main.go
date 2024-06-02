package main

import (
	"log"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/valyala/fasthttp"
)

func main() {
	// Create a new RateLimitedClient with 100 requests per second
	rlclient := client.NewRateLimitedClient(100, 1, &fasthttp.Client{})

	queryParams := map[string]string{
		"queryParam1": "value1",
		"queryParam2": "value2",
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	body := []byte(`{"key": "value"}`)

	// use fhc to make a GET request
	resp, err := client.Post(rlclient, "https://example.com", queryParams, headers, body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.StatusCode)
	log.Println(resp.Headers)
	log.Println(string(resp.Body))
	log.Println(resp.CurlCommand)
}
