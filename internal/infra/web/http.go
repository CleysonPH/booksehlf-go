package web

import "net/url"

type Headers struct {
	ContentType string
}

type HttpResponse struct {
	StatusCode int
	Body       []byte
	Headers    Headers
}

func NewHttpResponse(statusCode int, body []byte, headers Headers) *HttpResponse {
	return &HttpResponse{
		StatusCode: statusCode,
		Body:       body,
		Headers:    headers,
	}
}

type HttpRequest struct {
	QueryParams url.Values
	URLParams   map[string]string
	Body        []byte
}

func NewHttpRequest(
	queryParams url.Values,
	urlParams map[string]string,
	body []byte,
) *HttpRequest {
	return &HttpRequest{
		QueryParams: queryParams,
		URLParams:   urlParams,
		Body:        body,
	}
}
