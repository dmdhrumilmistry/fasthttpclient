# FastHttpClient

[![License](https://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/dmdhrumilmistry/fasthttpclient/blob/master/LICENSE)

FastHttpClient is a lightweight and high-performance HTTP client for sending requests.

## Features

- Simple and intuitive API
- Fast and efficient with optimized memory usage
- Supports various HTTP methods (GET, POST, PUT, DELETE, etc.)
- Use as Package
- Client with Rate Limiting option
- Comprehensive benchmark and unit tests

## Performance

FastHttpClient has been optimized for both performance and memory efficiency:
- Reduced memory allocations in URI generation and header processing
- Pre-allocated buffers to minimize reallocations
- Proper memory management with fasthttp object pools
- Comprehensive benchmarks included

For detailed information about performance optimizations, see [OPTIMIZATIONS.md](./OPTIMIZATIONS.md).

## Usage

- Add project to dependency

    ```bash
    go get github.com/dmdhrumilmistry/fasthttpclient
    ```

- Without Rate Limit

    ```go
    package main

    import (
        "log"

        "github.com/dmdhrumilmistry/fasthttpclient/client"
        "github.com/valyala/fasthttp"
    )

    func main() {
        // Create a new FHClient without any rate limit
        fhc := client.NewFHClient(&fasthttp.Client{})

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

- View all [examples](./examples/)
