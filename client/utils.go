package client

import (
	"bytes"
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
