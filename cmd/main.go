package main

import (
	"encoding/json"
	"log"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
)

func main() {
	// clients
	fastHttpClient := client.NewFHClient()
	// fastHttpClient := client.NewRateLimitedClient(10, 1)

	headers := map[string]string{
		"User-Agent":   "fasthttpclient",
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	queryParams := map[string]string{
		"format": "json",
	}

	bodyParams := make(map[string]interface{})
	bodyParams["name"] = "test"
	bodyParams["email"] = "test@example.com"
	bodyParams["listening"] = 100

	body, err := json.Marshal(bodyParams)
	if err != nil {
		log.Fatalln(err)
	}

	// resp, err := client.Get(fastHttpClient,"https://ipinfo.io", queryParams, headers)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	resp, err := client.Post(fastHttpClient, "http://localhost:8002/api/v1/forms/unqualified", queryParams, headers, body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.StatusCode)
	log.Println(resp.Headers)
	log.Println(string(resp.Body))
}
