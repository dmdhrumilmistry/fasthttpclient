package client

import (
	"time"

	"github.com/dmdhrumilmistry/fasthttpclient/config"
	"github.com/valyala/fasthttp"
	"golang.org/x/time/rate"
)

type FHClient struct {
	Client *fasthttp.Client
}

func NewFHClient() *FHClient {
	client := &fasthttp.Client{
		ReadTimeout:              config.Cfg.ReadTimeout,
		WriteTimeout:             config.Cfg.WriteTimeout,
		MaxIdleConnDuration:      config.Cfg.MaxIdleConn,
		NoDefaultUserAgentHeader: true, // Disable default User-Agent fasthttp header
		// increase DNS cache TTL to an hour
		Dial: (&fasthttp.TCPDialer{
			Concurrency:      4096,
			DNSCacheDuration: time.Hour,
		}).Dial,
	}

	return &FHClient{
		Client: client,
	}
}

type RateLimitedClient struct {
	Requests    int
	PerSeconds  int
	RateLimited *rate.Limiter
	FHClient    *FHClient
}

func NewRateLimitedClient(requests int, perSeconds int) *RateLimitedClient {
	return &RateLimitedClient{
		Requests:    requests,
		PerSeconds:  perSeconds,
		RateLimited: rate.NewLimiter(rate.Every(time.Second*time.Duration(perSeconds)), requests),
		FHClient:    NewFHClient(),
	}
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

func NewResponse(statusCode int, headers map[string]string, body []byte) *Response {
	return &Response{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}
}

type ClientInterface interface {
	Do(uri string, method string, queryParams any, headers any, body any) (*Response, error)
}
