package main

import (
	"log"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
)

func main() {
	headers := map[string]string{
		"User-Agent": "fasthttpclient",
		"Accept":     "application/json",
	}

	queryParams := map[string]string{
		"accept": "json",
	}

	resp, err := client.Get("https://ipinfo.io", queryParams, headers)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp.StatusCode)
	log.Println(resp.Headers)
	log.Println(string(resp.Body))
}
