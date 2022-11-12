package factory

import (
	"github.com/cleysonph/bookshelf-go/internal/application/gateway"
	"github.com/cleysonph/bookshelf-go/internal/application/usecase"
	"github.com/cleysonph/bookshelf-go/internal/infra/controller"
	gatewayImpl "github.com/cleysonph/bookshelf-go/internal/infra/gateway"
)

var bookGateway gateway.BookGateway

func BookGateway() gateway.BookGateway {
	if bookGateway == nil {
		bookGateway = gatewayImpl.NewBookInMemoryGateway()
	}
	return bookGateway
}

func GetBooksUseCase() *usecase.GetBooksUseCase {
	return usecase.NewGetBooksUseCase(BookGateway())
}

func GetBooksWebController() controller.WebController {
	return controller.NewGetBooksWebController(GetBooksUseCase())
}
