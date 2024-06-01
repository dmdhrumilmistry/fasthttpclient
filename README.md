# FastHttpClient

[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/dmdhrumilmistry/fasthttpclient/blob/master/LICENSE)

FastHttpClient is a lightweight and high-performance HTTP client for sending requests.

## Features

- Simple and intuitive API
- Fast and efficient
- Supports various HTTP methods (GET, POST, PUT, DELETE, etc.)
- Use as Package
- Client with Rate Limiting option

## Usage

- Add project to dependency

    ```bash
    go get github.com/dmdhrumilmistry/fasthttpclient
    ```

- configure timeouts in `.env` file

    ```txt
    FHC_READ_TIMEOUT=5s
    FHC_WRITE_TIMEOUT=5s
    FHC_MAX_IDLE_CONN=60s
    ```

    **OR**

    by exporting env variables

    ```bash
    export FHC_READ_TIMEOUT=5s
    export FHC_WRITE_TIMEOUT=5s
    export FHC_MAX_IDLE_CONN=60s
    ```

- Without Rate Limit

    ```go
    package main

    import (
        "log"

        "github.com/dmdhrumilmistry/fasthttpclient/client"
    )

    func main() {
        // Create a new FHClient without any rate limit
        fhc := client.NewFHClient()

        queryParams := map[string]string{
            "queryParam1": "value1",
            "queryParam2": "value2",
        }

        // use fhc to make a GET request
        resp, err := client.Get(fhc, "https://example.com", queryParams, nil)
        if err != nil {
            log.Fatalln(err)
        }

        log.Println(resp.StatusCode)
        log.Println(resp.Headers)
        log.Println(string(resp.Body))
        log.Println(resp.CurlCommand)
    }
    ```

- Using Rate Limit 100 requests/1sec

    ```go
    package main

    import (
        "log"

        "github.com/dmdhrumilmistry/fasthttpclient/client"
    )

    func main() {
        // Create a new RateLimitedClient with 100 requests per second
        rlclient := client.NewRateLimitedClient(100, 1)

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
    ```
