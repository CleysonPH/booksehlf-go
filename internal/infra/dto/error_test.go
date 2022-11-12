package dto

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrorResponse_WhenErrorIsNil_ShouldReturnNil(t *testing.T) {
	var err error
	response := NewErrorResponse(err, http.StatusInternalServerError)
	assert.Nil(t, response)
}

func TestNewErrorResponse_WhenErrorIsNotNil_ShouldReturnErrorResponse(t *testing.T) {
	err := errors.New("error")
	response := NewErrorResponse(err, http.StatusInternalServerError)
	assert.NotNil(t, response)
}
