package controller

import (
	"encoding/json"
	"net/http"

	"github.com/cleysonph/bookshelf-go/internal/application"
	"github.com/cleysonph/bookshelf-go/internal/application/validator"
	"github.com/cleysonph/bookshelf-go/internal/domain"
	"github.com/cleysonph/bookshelf-go/internal/infra/dto"
	"github.com/cleysonph/bookshelf-go/internal/infra/web"
)

type WebController interface {
	Execute(request *web.HttpRequest) *web.HttpResponse
}

var defaultHeaders = web.Headers{
	ContentType: "application/json",
}

func newJsonResponse(statusCode int, body interface{}) *web.HttpResponse {
	json, err := json.Marshal(body)
	if err != nil {
		return handleErrorResponse(err)
	}
	return web.NewHttpResponse(statusCode, json, defaultHeaders)
}

func handleErrorResponse(err error) *web.HttpResponse {
	switch t := err.(type) {
	case *application.BookNotFoundError:
		e := dto.NewErrorResponse(t, http.StatusNotFound)
		return newJsonResponse(e.Status, e)
	case *application.ApplicationError:
		e := dto.NewErrorResponse(t, http.StatusInternalServerError)
		return newJsonResponse(e.Status, e)
	case *validator.ValidationError:
		e := dto.NewValidationErrorResponse(t, http.StatusBadRequest, t.Errors)
		return newJsonResponse(e.Status, e)
	case *json.SyntaxError:
		e := dto.NewErrorResponse(t, http.StatusBadRequest)
		return newJsonResponse(e.Status, e)
	case *domain.DomainError:
		e := dto.NewErrorResponse(t, http.StatusBadRequest)
		return newJsonResponse(e.Status, e)
	default:
		e := dto.NewErrorResponse(err, http.StatusInternalServerError)
		return newJsonResponse(e.Status, e)
	}
}
