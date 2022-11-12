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
}

func NewHttpRequest(queryParams url.Values) *HttpRequest {
	return &HttpRequest{
		QueryParams: queryParams,
	}
}
