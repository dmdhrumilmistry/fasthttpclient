package client

import (
	"time"

	"github.com/dmdhrumilmistry/fasthttpclient/config"
	"github.com/valyala/fasthttp"
)

var client = &fasthttp.Client{
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

func Get(uri string, queryParams any, headers any) (*Response, error) {
	// generate request uri
	uri, err := SetQueryParamsInURI(queryParams, uri)
	if err != nil {
		return nil, err
	}

	// acquire resources for request
	req := fasthttp.AcquireRequest()

	// configure uri and method
	req.SetRequestURI(uri)
	req.Header.SetMethod(fasthttp.MethodGet)

	// set headers
	if err := SetHeadersInRequest(headers, req); err != nil {
		return nil, err
	}

	// acquire response
	resp := fasthttp.AcquireResponse()
	err = client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	// release request resources
	fasthttp.ReleaseRequest(req)

	// release response after using it
	body := resp.Body()
	respHeaders := GetResponseHeaders(resp)
	statusCode := resp.StatusCode()

	// release response body resources
	fasthttp.ReleaseResponse(resp)

	return NewResponse(statusCode, respHeaders, body), nil
}
