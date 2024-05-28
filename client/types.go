package client

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       []byte
}

func NewResponse(statusCode int, headers map[string]string, body []byte) *Response {
	return &Response{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}
}
