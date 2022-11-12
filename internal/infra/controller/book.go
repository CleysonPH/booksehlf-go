package controller

import (
	"net/http"

	"github.com/cleysonph/bookshelf-go/internal/application/usecase"
	"github.com/cleysonph/bookshelf-go/internal/infra/dto"
	"github.com/cleysonph/bookshelf-go/internal/infra/web"
)

type GetBooksWebController struct {
	getBooksUseCase *usecase.GetBooksUseCase
}

// Execute implements WebController
func (g *GetBooksWebController) Execute(request *web.HttpRequest) *web.HttpResponse {
	q := request.QueryParams.Get("q")
	books, err := g.getBooksUseCase.Execute(q)
	if err != nil {
		return handleErrorResponse(err)
	}
	body := make([]dto.BookSummaryResponse, len(books))
	for i, book := range books {
		b := dto.BookSummaryResponse{}
		b.FromDomain(book)
		body[i] = b
	}
	return newJsonResponse(http.StatusOK, body)
}

func NewGetBooksWebController(getBooksUseCase *usecase.GetBooksUseCase) WebController {
	return &GetBooksWebController{getBooksUseCase: getBooksUseCase}
}
