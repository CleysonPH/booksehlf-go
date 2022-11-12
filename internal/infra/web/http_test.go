package web

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHttpResponse_ShouldReturnHttpResponse(t *testing.T) {
	expectedStatusCode := 200
	expectedBody := []byte("body")
	expectedHeaders := Headers{ContentType: "application/json"}
	response := NewHttpResponse(expectedStatusCode, expectedBody, Headers{ContentType: "application/json"})
	assert.Equal(t, expectedStatusCode, response.StatusCode)
	assert.Equal(t, expectedBody, response.Body)
	assert.Equal(t, expectedHeaders, response.Headers)
}

func TestNewHttpRequest_ShouldReturnHttpRequest(t *testing.T) {
	expectedQueryParams := url.Values{
		"key": {"value"},
	}
	expectedURLParams := map[string]string{
		"key": "value",
	}
	request := NewHttpRequest(expectedQueryParams, expectedURLParams)
	assert.Equal(t, expectedQueryParams, request.QueryParams)
	assert.Equal(t, expectedURLParams, request.URLParams)
}
