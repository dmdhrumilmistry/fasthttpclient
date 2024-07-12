package client

import (
	"time"

	"github.com/valyala/fasthttp"
	"golang.org/x/time/rate"
)

type FHClient struct {
	Client *fasthttp.Client
}

func NewFHClient(fasthttp *fasthttp.Client) *FHClient {
	return &FHClient{
		Client: fasthttp,
	}
}

type RateLimitedClient struct {
	Requests    int
	PerSeconds  int
	RateLimited *rate.Limiter
	FHClient    *FHClient
}

func NewRateLimitedClient(requests int, perSeconds int, fasthttp *fasthttp.Client) *RateLimitedClient {
	return &RateLimitedClient{
		Requests:    requests,
		PerSeconds:  perSeconds,
		RateLimited: rate.NewLimiter(rate.Every(time.Second*time.Duration(perSeconds)), requests),
		FHClient:    NewFHClient(fasthttp),
	}
}

type Response struct {
	StatusCode  int
	Headers     map[string]string
	Body        []byte
	CurlCommand string
	TimeElapsed time.Duration
}

func NewResponse(statusCode int, headers map[string]string, body []byte, curlCmd string, duration time.Duration) *Response {
	return &Response{
		StatusCode:  statusCode,
		Headers:     headers,
		Body:        body,
		CurlCommand: curlCmd,
		TimeElapsed: duration,
	}
}

type Request struct {
	Uri         string
	Method      string
	QueryParams any
	Headers     any
	Body        any
}

func NewRequest(uri string, method string, queryParams any, headers any, body any) *Request {
	return &Request{
		Uri:         uri,
		Method:      method,
		QueryParams: queryParams,
		Headers:     headers,
		Body:        body,
	}
}

type ConcurrentResponse struct {
	Response *Response
	Error    error
}

func NewConcurrentResponse(response *Response, err error) *ConcurrentResponse {
	return &ConcurrentResponse{
		Response: response,
		Error:    err,
	}
}

type ClientInterface interface {
	Do(uri string, method string, queryParams any, headers any, body any) (*Response, error)
}
