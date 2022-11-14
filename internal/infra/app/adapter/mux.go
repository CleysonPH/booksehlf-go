package adapter

import (
	"io"
	"net/http"

	"github.com/cleysonph/bookshelf-go/internal/infra/controller"
	"github.com/cleysonph/bookshelf-go/internal/infra/web"
	"github.com/gorilla/mux"
)

func MuxAdapt(webController controller.WebController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			body = []byte{}
		}
		req := web.NewHttpRequest(r.URL.Query(), mux.Vars(r), body)
		res := webController.Execute(req)
		w.WriteHeader(res.StatusCode)
		w.Header().Set("Content-Type", res.Headers.ContentType)
		w.Write(res.Body)
	}
}
