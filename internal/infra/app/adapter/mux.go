package adapter

import (
	"net/http"

	"github.com/cleysonph/bookshelf-go/internal/infra/controller"
	"github.com/cleysonph/bookshelf-go/internal/infra/web"
	"github.com/gorilla/mux"
)

func MuxAdapt(webController controller.WebController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := web.NewHttpRequest(r.URL.Query(), mux.Vars(r))
		res := webController.Execute(req)
		w.WriteHeader(res.StatusCode)
		w.Header().Set("Content-Type", res.Headers.ContentType)
		w.Write(res.Body)
	}
}
