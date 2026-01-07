package client_test

import (
	"strings"
	"testing"
	"time"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/valyala/fasthttp"
)

// Test GenerateURI with various inputs
func TestGenerateURI(t *testing.T) {
	tests := []struct {
		name        string
		uri         string
		queryParams map[string]string
		wantContain []string
	}{
		{
			name: "Empty query params",
			uri:  "https://example.com",
			queryParams: map[string]string{},
			wantContain: []string{"https://example.com"},
		},
		{
			name: "Single query param",
			uri:  "https://example.com",
			queryParams: map[string]string{
				"key": "value",
			},
			wantContain: []string{"https://example.com?", "key=value"},
		},
		{
			name: "Multiple query params",
			uri:  "https://example.com/api",
			queryParams: map[string]string{
				"param1": "value1",
				"param2": "value2",
			},
			wantContain: []string{"https://example.com/api?", "param1=value1", "param2=value2"},
		},
		{
			name: "Special characters in params",
			uri:  "https://example.com",
			queryParams: map[string]string{
				"email": "test@example.com",
			},
			wantContain: []string{"https://example.com?", "email="},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := client.GenerateURI(tt.uri, tt.queryParams)
			for _, want := range tt.wantContain {
				if !strings.Contains(got, want) {
					t.Errorf("GenerateURI() = %v, want to contain %v", got, want)
				}
			}
		})
	}
}

// Test SetQueryParamsInURI
func TestSetQueryParamsInURI(t *testing.T) {
	tests := []struct {
		name        string
		uri         string
		queryParams interface{}
		wantErr     bool
		wantContain string
	}{
		{
			name: "Valid map[string]string",
			uri:  "https://example.com",
			queryParams: map[string]string{
				"key": "value",
			},
			wantErr:     false,
			wantContain: "key=value",
		},
		{
			name:        "Nil query params",
			uri:         "https://example.com",
			queryParams: nil,
			wantErr:     false,
			wantContain: "https://example.com",
		},
		{
			name:        "Invalid query params type",
			uri:         "https://example.com",
			queryParams: "invalid",
			wantErr:     true,
			wantContain: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.SetQueryParamsInURI(tt.queryParams, tt.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetQueryParamsInURI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !strings.Contains(got, tt.wantContain) {
				t.Errorf("SetQueryParamsInURI() = %v, want to contain %v", got, tt.wantContain)
			}
		})
	}
}

// Test GetResponseHeaders
func TestGetResponseHeaders(t *testing.T) {
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("Content-Length", "1234")
	resp.Header.Set("Server", "fasthttp")

	headers := client.GetResponseHeaders(resp)

	if len(headers) != 3 {
		t.Errorf("Expected 3 headers, got %d", len(headers))
	}

	if headers["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type=application/json, got %v", headers["Content-Type"])
	}

	if headers["Content-Length"] != "1234" {
		t.Errorf("Expected Content-Length=1234, got %v", headers["Content-Length"])
	}

	if headers["Server"] != "fasthttp" {
		t.Errorf("Expected Server=fasthttp, got %v", headers["Server"])
	}
}

// Test SetHeaders
func TestSetHeaders(t *testing.T) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	headers := map[string]string{
		"Content-Type":  "application/json",
		"User-Agent":    "test-agent",
		"Authorization": "Bearer token123",
	}

	client.SetHeaders(req, headers)

	if string(req.Header.Peek("Content-Type")) != "application/json" {
		t.Errorf("Expected Content-Type=application/json, got %v", string(req.Header.Peek("Content-Type")))
	}

	if string(req.Header.Peek("User-Agent")) != "test-agent" {
		t.Errorf("Expected User-Agent=test-agent, got %v", string(req.Header.Peek("User-Agent")))
	}

	if string(req.Header.Peek("Authorization")) != "Bearer token123" {
		t.Errorf("Expected Authorization=Bearer token123, got %v", string(req.Header.Peek("Authorization")))
	}
}

// Test SetHeadersInRequest
func TestSetHeadersInRequest(t *testing.T) {
	tests := []struct {
		name    string
		headers interface{}
		wantErr bool
	}{
		{
			name: "Valid headers map",
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			wantErr: false,
		},
		{
			name:    "Nil headers",
			headers: nil,
			wantErr: false,
		},
		{
			name:    "Invalid headers type",
			headers: "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := fasthttp.AcquireRequest()
			defer fasthttp.ReleaseRequest(req)

			err := client.SetHeadersInRequest(tt.headers, req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetHeadersInRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// Test SetRequestBody
func TestSetRequestBody(t *testing.T) {
	tests := []struct {
		name     string
		body     interface{}
		wantErr  bool
		wantBody string
	}{
		{
			name:     "Valid byte slice",
			body:     []byte("test body"),
			wantErr:  false,
			wantBody: "test body",
		},
		{
			name:     "Nil body",
			body:     nil,
			wantErr:  false,
			wantBody: "",
		},
		{
			name:     "Invalid body type",
			body:     "invalid string",
			wantErr:  true,
			wantBody: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := fasthttp.AcquireRequest()
			defer fasthttp.ReleaseRequest(req)

			err := client.SetRequestBody(tt.body, req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetRequestBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && string(req.Body()) != tt.wantBody {
				t.Errorf("SetRequestBody() body = %v, want %v", string(req.Body()), tt.wantBody)
			}
		})
	}
}

// Test NewResponse
func TestNewResponse(t *testing.T) {
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	body := []byte(`{"status": "ok"}`)
	curlCmd := "curl -X GET https://example.com"
	duration := time.Duration(100)

	resp := client.NewResponse(200, headers, body, curlCmd, duration)

	if resp.StatusCode != 200 {
		t.Errorf("Expected StatusCode=200, got %d", resp.StatusCode)
	}

	if resp.Headers["Content-Type"] != "application/json" {
		t.Errorf("Expected Content-Type=application/json, got %v", resp.Headers["Content-Type"])
	}

	if string(resp.Body) != `{"status": "ok"}` {
		t.Errorf("Expected body={\"status\": \"ok\"}, got %v", string(resp.Body))
	}

	if resp.CurlCommand != curlCmd {
		t.Errorf("Expected CurlCommand=%v, got %v", curlCmd, resp.CurlCommand)
	}

	if resp.TimeElapsed != duration {
		t.Errorf("Expected TimeElapsed=%v, got %v", duration, resp.TimeElapsed)
	}
}

// Test NewConcurrentResponse
func TestNewConcurrentResponse(t *testing.T) {
	resp := &client.Response{
		StatusCode: 200,
		Body:       []byte("test"),
	}

	concResp := client.NewConcurrentResponse(resp, nil)

	if concResp.Response != resp {
		t.Errorf("Expected Response=%v, got %v", resp, concResp.Response)
	}

	if concResp.Error != nil {
		t.Errorf("Expected Error=nil, got %v", concResp.Error)
	}
}

// Test NewRequest
func TestNewRequest(t *testing.T) {
	uri := "https://example.com"
	method := "GET"
	queryParams := map[string]string{"key": "value"}
	headers := map[string]string{"Content-Type": "application/json"}
	body := []byte("test body")

	req := client.NewRequest(uri, method, queryParams, headers, body)

	if req.Uri != uri {
		t.Errorf("Expected Uri=%v, got %v", uri, req.Uri)
	}

	if req.Method != method {
		t.Errorf("Expected Method=%v, got %v", method, req.Method)
	}
}
