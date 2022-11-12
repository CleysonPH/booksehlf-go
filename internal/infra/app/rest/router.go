package rest

import (
	"net/http"

	"github.com/cleysonph/bookshelf-go/internal/infra/app/adapter"
	"github.com/cleysonph/bookshelf-go/internal/infra/app/factory"
	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/books", adapter.MuxAdapt(factory.GetBooksWebController())).
		Methods(http.MethodGet)

	return router
}