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
		Timestamp: time.Now(),
		Status:    satusCode,
		Error:     http.StatusText(satusCode),
		Cause:     strings.Split(reflect.TypeOf(err).String(), ".")[1],
	}
}
