package client

import (
	"errors"
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
	queryParamsMap, ok := queryParams.(map[string]string)
	if !ok && queryParams != nil {
		return nil, errors.New("queryParams must be a map[string]string")
	} else {
		uri = GenerateURI(uri, queryParamsMap)
	}

	req := fasthttp.AcquireRequest()

	// configure uri and method
	req.SetRequestURI(uri)
	req.Header.SetMethod(fasthttp.MethodGet)

	// set headers
	headersMap, ok := headers.(map[string]string)
	if !ok && headers != nil {
		return nil, errors.New("headers must be a map[string]string")
	} else {
		SetHeaders(req, headersMap)
	}

	resp := fasthttp.AcquireResponse()
	err := client.Do(req, resp)
	if err != nil {
		return nil, err
	}
	fasthttp.ReleaseRequest(req)

	// release response after using it
	body := resp.Body()
	respHeaders := GetResponseHeaders(resp)
	statusCode := resp.StatusCode()

	fasthttp.ReleaseResponse(resp)

	return NewResponse(statusCode, respHeaders, body), nil
}
