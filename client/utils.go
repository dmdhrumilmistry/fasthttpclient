package client

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"

	"github.com/valyala/fasthttp"
)

func SetHeaders(req *fasthttp.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

func GetResponseHeaders(resp *fasthttp.Response) map[string]string {
	headers := make(map[string]string)
	resp.Header.VisitAll(func(key, value []byte) {
		headers[string(key)] = string(value)
	})
	return headers
}

func GenerateURI(uri string, queryParams map[string]string) string {
	var urlBuffer bytes.Buffer
	urlBuffer.WriteString(uri + "?")
	for key, value := range queryParams {
		fmt.Fprintf(&urlBuffer, "%s=%s&", url.QueryEscape(key), url.QueryEscape(value))
	}
	return urlBuffer.String()[:len(urlBuffer.String())-1] // Remove trailing "&"
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
