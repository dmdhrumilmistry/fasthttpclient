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
	StatusCode  int               `json:"status_code" yaml:"status_code"`
	Headers     map[string]string `json:"headers" yaml:"headers"`
	Body        []byte            `json:"body" yaml:"body"`
	CurlCommand string            `json:"curl_command" yaml:"curl_command"`
	TimeElapsed time.Duration     `json:"time_elapsed" yaml:"time_elapsed"`
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
	Uri         string      `json:"uri" yaml:"uri"`
	Method      string      `json:"method" yaml:"method"`
	QueryParams interface{} `json:"query_params" yaml:"query_params"`
	Headers     interface{} `json:"headers" yaml:"headers"`
	Body        interface{} `json:"body" yaml:"body"`
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
	Response *Response `json:"response" yaml:"response"`
	Error    error     `json:"error" yaml:"error"`
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
