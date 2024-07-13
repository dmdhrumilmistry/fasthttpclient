package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/valyala/fasthttp"
)

func main() {
	// create new client and requests holder
	fhc := client.NewFHClient(&fasthttp.Client{
		MaxConnsPerHost:          10000,
		ReadTimeout:              time.Second * 5,
		WriteTimeout:             time.Second * 5,
		MaxIdleConnDuration:      time.Second * 60,
		NoDefaultUserAgentHeader: true,
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	})
	requests := []*client.Request{}

	queryParams := map[string]string{
		"test": "true",
	}

	headers := map[string]string{
		"User-Agent": "fasthttpclient",
	}

	reqCount := 5
	// create requests
	for i := 0; i < reqCount; i++ {
		requests = append(requests, client.NewRequest("https://example.com", fasthttp.MethodGet, queryParams, headers, nil))
	}

	// make concurrent requests
	responses := client.MakeConcurrentRequests(requests, fhc)
	log.Printf("\n%d Requests Completed\n", len(responses))

	jsondata, err := json.Marshal(responses)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(jsondata))
}
