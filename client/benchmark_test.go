package client_test

import (
	"testing"

	"github.com/dmdhrumilmistry/fasthttpclient/client"
	"github.com/valyala/fasthttp"
)

// Benchmark for GenerateURI function
func BenchmarkGenerateURI(b *testing.B) {
	uri := "https://example.com/api"
	queryParams := map[string]string{
		"param1": "value1",
		"param2": "value2",
		"param3": "value3",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.GenerateURI(uri, queryParams)
	}
}

// Benchmark for GenerateURI with many parameters
func BenchmarkGenerateURIMany(b *testing.B) {
	uri := "https://example.com/api"
	queryParams := map[string]string{
		"param1":  "value1",
		"param2":  "value2",
		"param3":  "value3",
		"param4":  "value4",
		"param5":  "value5",
		"param6":  "value6",
		"param7":  "value7",
		"param8":  "value8",
		"param9":  "value9",
		"param10": "value10",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.GenerateURI(uri, queryParams)
	}
}

// Benchmark for GenerateURI with empty params
func BenchmarkGenerateURIEmpty(b *testing.B) {
	uri := "https://example.com/api"
	queryParams := map[string]string{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.GenerateURI(uri, queryParams)
	}
}

// Benchmark for GetResponseHeaders
func BenchmarkGetResponseHeaders(b *testing.B) {
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Add some headers
	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("Content-Length", "1234")
	resp.Header.Set("Server", "fasthttp")
	resp.Header.Set("Date", "Mon, 01 Jan 2024 00:00:00 GMT")
	resp.Header.Set("X-Custom-Header", "custom-value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.GetResponseHeaders(resp)
	}
}

// Benchmark for SetHeaders
func BenchmarkSetHeaders(b *testing.B) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	headers := map[string]string{
		"Content-Type":  "application/json",
		"User-Agent":    "fasthttpclient",
		"Accept":        "application/json",
		"Authorization": "Bearer token123",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.SetHeaders(req, headers)
	}
}

// Benchmark for SetQueryParamsInURI
func BenchmarkSetQueryParamsInURI(b *testing.B) {
	uri := "https://example.com/api"
	queryParams := map[string]string{
		"param1": "value1",
		"param2": "value2",
		"param3": "value3",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = client.SetQueryParamsInURI(queryParams, uri)
	}
}

// Benchmark for NewResponse
func BenchmarkNewResponse(b *testing.B) {
	headers := map[string]string{
		"Content-Type": "application/json",
		"Server":       "fasthttp",
	}
	body := []byte(`{"status": "ok", "message": "test"}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.NewResponse(200, headers, body, "curl command", 100)
	}
}

// Benchmark for memory allocations in concurrent scenario
func BenchmarkConcurrentResponseAllocation(b *testing.B) {
	resp := &client.Response{
		StatusCode:  200,
		Headers:     map[string]string{"Content-Type": "application/json"},
		Body:        []byte(`{"test": "data"}`),
		CurlCommand: "curl test",
		TimeElapsed: 100,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = client.NewConcurrentResponse(resp, nil)
	}
}
