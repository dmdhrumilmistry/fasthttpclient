package client

import "github.com/valyala/fasthttp"

func Connect(client ClientInterface, uri string, queryParams any, headers any) (*Response, error) {
	return client.Do(uri, fasthttp.MethodPost, queryParams, headers, nil)
}

func Delete(client ClientInterface, uri string, queryParams any, headers any, body any) (*Response, error) {
	return client.Do(uri, fasthttp.MethodPost, queryParams, headers, body)
}

func Head(client ClientInterface, uri string, queryParams any, headers any) (*Response, error) {
	return client.Do(uri, fasthttp.MethodPost, queryParams, headers, nil)
}

func Get(client ClientInterface, uri string, queryParams any, headers any) (*Response, error) {
	return client.Do(uri, fasthttp.MethodGet, queryParams, headers, nil)
}

func Options(client ClientInterface, uri string, queryParams any, headers any) (*Response, error) {
	return client.Do(uri, fasthttp.MethodPost, queryParams, headers, nil)
}

func Post(client ClientInterface, uri string, queryParams any, headers any, body any) (*Response, error) {
	return client.Do(uri, fasthttp.MethodPost, queryParams, headers, body)
}

func Put(client ClientInterface, uri string, queryParams any, headers any, body any) (*Response, error) {
	return client.Do(uri, fasthttp.MethodPut, queryParams, headers, body)
}

func Patch(client ClientInterface, uri string, queryParams any, headers any, body any) (*Response, error) {
	return client.Do(uri, fasthttp.MethodPost, queryParams, headers, body)
}

func Trace(client ClientInterface, uri string, queryParams any, headers any) (*Response, error) {
	return client.Do(uri, fasthttp.MethodPost, queryParams, headers, nil)
}
