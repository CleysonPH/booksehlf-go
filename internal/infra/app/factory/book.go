package factory

import (
	"github.com/cleysonph/bookshelf-go/internal/application/gateway"
	"github.com/cleysonph/bookshelf-go/internal/application/usecase"
	"github.com/cleysonph/bookshelf-go/internal/application/validator"
	"github.com/cleysonph/bookshelf-go/internal/infra/controller"
	gatewayImpl "github.com/cleysonph/bookshelf-go/internal/infra/gateway"
	validatorImpl "github.com/cleysonph/bookshelf-go/internal/infra/validator"
)

var bookGateway gateway.BookGateway

func BookGateway() gateway.BookGateway {
	if bookGateway == nil {
		bookGateway = gatewayImpl.NewBookMySQLGateway(gatewayImpl.DB())
	}
	return bookGateway
}

func CreateBookValidator() validator.CreateBookValidator {
	return validatorImpl.NewCreateBookValidator()
}

func GetBooksUseCase() *usecase.GetBooksUseCase {
	return usecase.NewGetBooksUseCase(BookGateway())
}

func GetBookUseCase() *usecase.GetBookUseCase {
	return usecase.NewGetBookUseCase(BookGateway())
}

func DeleteBookUseCase() *usecase.DeleteBookUseCase {
	return usecase.NewDeleteBookUseCase(BookGateway())
}

func CreateBookUseCase() *usecase.CreateBookUseCase {
	return usecase.NewCreateBookUseCase(BookGateway(), CreateBookValidator())
}

func GetBooksWebController() controller.WebController {
	return controller.NewGetBooksWebController(GetBooksUseCase())
}

func GetBookWebController() controller.WebController {
	return controller.NewGetBookWebController(GetBookUseCase())
}

func DeleteBookWebController() controller.WebController {
	return controller.NewDeleteBookWebController(DeleteBookUseCase())
}

func CreateBookWebController() controller.WebController {
	return controller.NewCreateBookWebController(CreateBookUseCase())
}
