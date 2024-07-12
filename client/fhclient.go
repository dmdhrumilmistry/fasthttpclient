package client

import (
	"time"

	"github.com/li-jin-gou/http2curl"
	"github.com/valyala/fasthttp"
)

func (c *FHClient) Do(uri string, method string, queryParams any, headers any, reqBody any) (*Response, error) {
	// start timer
	now := time.Now()

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

	// get curl command
	curlCmd := "error generating curl command"
	curlCmdObj, err := http2curl.GetCurlCommandFastHttp(req)
	if err == nil {
		curlCmd = curlCmdObj.String()
	}

	// acquire response
	resp := fasthttp.AcquireResponse()
	err = c.Client.Do(req, resp)
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

	// stop timer
	elapsed := time.Since(now)
	return NewResponse(statusCode, respHeaders, body, curlCmd, elapsed), nil
}
