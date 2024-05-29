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

func Do(uri string, method string, queryParams any, headers any, reqBody any) (*Response, error) {
	// generate request uri
	uri, err := SetQueryParamsInURI(queryParams, uri)
	if err != nil {
		return nil, err
	}

	// acquire resources for request
	req := fasthttp.AcquireRequest()

	// configure uri and method
	req.SetRequestURI(uri)
	req.Header.SetMethod(method)

	// set headers
	if err := SetHeadersInRequest(headers, req); err != nil {
		return nil, err
	}

	// set body if valid
	if err := SetRequestBody(reqBody, req); err != nil {
		return nil, err
	}

	// acquire response
	resp := fasthttp.AcquireResponse()
	err = client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	// release resources after use
	fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	// release response after using it
	body := resp.Body()
	respHeaders := GetResponseHeaders(resp)
	statusCode := resp.StatusCode()

	return NewResponse(statusCode, respHeaders, body), nil
}

func Connect(uri string, queryParams any, headers any) (*Response, error) {
	return Do(uri, fasthttp.MethodPost, queryParams, headers, nil)
}

func Delete(uri string, queryParams any, headers any, body any) (*Response, error) {
	return Do(uri, fasthttp.MethodPost, queryParams, headers, body)
}

func Head(uri string, queryParams any, headers any) (*Response, error) {
	return Do(uri, fasthttp.MethodPost, queryParams, headers, nil)
}

func Get(uri string, queryParams any, headers any) (*Response, error) {
	return Do(uri, fasthttp.MethodGet, queryParams, headers, nil)
}

func Options(uri string, queryParams any, headers any) (*Response, error) {
	return Do(uri, fasthttp.MethodPost, queryParams, headers, nil)
}

func Post(uri string, queryParams any, headers any, body any) (*Response, error) {
	return Do(uri, fasthttp.MethodPost, queryParams, headers, body)
}

func Put(uri string, queryParams any, headers any, body any) (*Response, error) {
	return Do(uri, fasthttp.MethodPut, queryParams, headers, body)
}

func Patch(uri string, queryParams any, headers any, body any) (*Response, error) {
	return Do(uri, fasthttp.MethodPost, queryParams, headers, body)
}

func Trace(uri string, queryParams any, headers any) (*Response, error) {
	return Do(uri, fasthttp.MethodPost, queryParams, headers, nil)
}
