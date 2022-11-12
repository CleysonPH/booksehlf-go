package controller

import (
	"errors"
	"strings"
	"testing"

	"github.com/cleysonph/bookshelf-go/internal/application"
	"github.com/cleysonph/bookshelf-go/internal/infra/web"
	"github.com/stretchr/testify/assert"
)

func TestNewJsonResponse_ShouldReturnJsonResponse(t *testing.T) {
	expectedStatusCode := 200
	expectedBody := []byte("{\"message\":\"Hello World\"}")
	expectedHeaders := web.Headers{ContentType: "application/json"}
	response := newJsonResponse(expectedStatusCode, map[string]string{"message": "Hello World"})
	assert.Equal(t, expectedStatusCode, response.StatusCode)
	assert.Equal(t, expectedBody, response.Body)
	assert.Equal(t, expectedHeaders, response.Headers)
}

func TestNewJsonResponse_WhenJsonMarshalFail_ShouldReturnJsonResponse(t *testing.T) {
	expectedStatusCode := 500
	expectedHeaders := web.Headers{ContentType: "application/json"}
	response := newJsonResponse(expectedStatusCode, func() {})
	assert.Equal(t, expectedStatusCode, response.StatusCode)
	assert.True(t, strings.Contains(string(response.Body), `"message":"json: unsupported type: func()"`))
	assert.Equal(t, expectedHeaders, response.Headers)
}

func TestHandleErrorResponse_WhenErrorIsApplicationError_ShouldReturnJsonResponse(t *testing.T) {
	expectedStatusCode := 500
	expectedHeaders := web.Headers{ContentType: "application/json"}
	applicationError := application.NewApplicationError(errors.New("error"), "error")
	response := handleErrorResponse(applicationError)
	assert.Equal(t, expectedStatusCode, response.StatusCode)
	assert.True(t, strings.Contains(string(response.Body), `"message":"error"`))
	assert.Equal(t, expectedHeaders, response.Headers)
}
