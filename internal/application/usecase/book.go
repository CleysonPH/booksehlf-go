package usecase

import (
	"github.com/cleysonph/bookshelf-go/internal/application/gateway"
	"github.com/cleysonph/bookshelf-go/internal/domain"
)

type GetBooksUseCase struct {
	bookGateway gateway.BookGateway
}

func NewGetBooksUseCase(bookGateway gateway.BookGateway) *GetBooksUseCase {
	return &GetBooksUseCase{bookGateway}
}

func (u *GetBooksUseCase) Execute(title string) ([]*domain.Book, error) {
	return u.bookGateway.FindAllByTitle(title)
}

type GetBookUseCase struct {
	bookGateway gateway.BookGateway
}

func NewGetBookUseCase(bookGateway gateway.BookGateway) *GetBookUseCase {
	return &GetBookUseCase{bookGateway}
}

func (u *GetBookUseCase) Execute(id string) (*domain.Book, error) {
	return u.bookGateway.FindById(id)
}

type DeleteBookUseCase struct {
	bookGateway gateway.BookGateway
}

func NewDeleteBookUseCase(bookGateway gateway.BookGateway) *DeleteBookUseCase {
	return &DeleteBookUseCase{bookGateway}
}

func (u *DeleteBookUseCase) Execute(id string) error {
	return u.bookGateway.DeleteById(id)
}
