package dto

import (
	"net/http"
	"reflect"
	"strings"
	"time"
)

type ErrorResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
	Cause     string    `json:"cause"`
}

func NewErrorResponse(err error, satusCode int) *ErrorResponse {
	if err == nil {
		return nil
	}
	return &ErrorResponse{
		Message:   err.Error(),
		Timestamp: time.Now().UTC(),
		Status:    satusCode,
		Error:     http.StatusText(satusCode),
		Cause:     strings.Split(reflect.TypeOf(err).String(), ".")[1],
	}
}

type ValidationErrorResponse struct {
	Message   string              `json:"message"`
	Timestamp time.Time           `json:"timestamp"`
	Status    int                 `json:"status"`
	Error     string              `json:"error"`
	Cause     string              `json:"cause"`
	Errors    map[string][]string `json:"errors"`
}

func NewValidationErrorResponse(err error, satusCode int, errors map[string][]string) *ValidationErrorResponse {
	if err == nil {
		return nil
	}
	return &ValidationErrorResponse{
		Message:   err.Error(),
		Timestamp: time.Now().UTC(),
		Status:    satusCode,
		Error:     http.StatusText(satusCode),
		Cause:     strings.Split(reflect.TypeOf(err).String(), ".")[1],
		Errors:    errors,
	}
}
