package client

import (
	"bytes"
	"errors"
	"net/url"

	"github.com/valyala/fasthttp"
)

func SetHeaders(req *fasthttp.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

func GetResponseHeaders(resp *fasthttp.Response) map[string]string {
	// Count headers first to pre-allocate map
	headerCount := 0
	resp.Header.VisitAll(func(key, value []byte) {
		headerCount++
	})
	
	headers := make(map[string]string, headerCount)
	resp.Header.VisitAll(func(key, value []byte) {
		headers[string(key)] = string(value)
	})
	return headers
}

func GenerateURI(uri string, queryParams map[string]string) string {
	if len(queryParams) == 0 {
		return uri
	}
	
	var urlBuffer bytes.Buffer
	urlBuffer.Grow(len(uri) + 100) // Pre-allocate to reduce reallocations
	urlBuffer.WriteString(uri)
	urlBuffer.WriteByte('?')
	
	first := true
	for key, value := range queryParams {
		if !first {
			urlBuffer.WriteByte('&')
		}
		first = false
		urlBuffer.WriteString(url.QueryEscape(key))
		urlBuffer.WriteByte('=')
		urlBuffer.WriteString(url.QueryEscape(value))
	}
	return urlBuffer.String()
}

// SetQueryParamsInURI sets query parameters in the URI
// returns uri string without any error
func SetQueryParamsInURI(queryParams interface{}, uri string) (string, error) {
	queryParamsMap, ok := queryParams.(map[string]string)
	if !ok && queryParams != nil {
		return uri, errors.New("queryParams must be a map[string]string")
	} else {
		uri = GenerateURI(uri, queryParamsMap)
	}

	return uri, nil
}

func SetHeadersInRequest(headers interface{}, req *fasthttp.Request) error {
	// set headers
	headersMap, ok := headers.(map[string]string)
	if !ok && headers != nil {
		return errors.New("headers must be a map[string]string")
	} else {
		SetHeaders(req, headersMap)
	}

	return nil
}

// sets body if present and returns nil as error.
// Function will also return nil if body == nil
func SetRequestBody(body interface{}, req *fasthttp.Request) error {
	if body == nil {
		return nil
	}

	bodyBytes, ok := body.([]byte)
	if !ok {
		return errors.New("body only supports []byte type")
	} else {
		req.SetBody(bodyBytes)
	}
	return nil
}
