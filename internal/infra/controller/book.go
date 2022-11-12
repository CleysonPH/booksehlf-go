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

type GetBookWebController struct {
	getBookUseCase *usecase.GetBookUseCase
}

// Execute implements WebController
func (g *GetBookWebController) Execute(request *web.HttpRequest) *web.HttpResponse {
	id := request.URLParams["bookId"]
	book, err := g.getBookUseCase.Execute(id)
	if err != nil {
		return handleErrorResponse(err)
	}
	body := dto.BookResponse{}
	body.FromDomain(book)
	return newJsonResponse(http.StatusOK, body)
}

func NewGetBookWebController(getBookUseCase *usecase.GetBookUseCase) WebController {
	return &GetBookWebController{getBookUseCase: getBookUseCase}
}
