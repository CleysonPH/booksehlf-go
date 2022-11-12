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

func (u *GetBookUseCase) Execute(id int64) (*domain.Book, error) {
	return u.bookGateway.FindById(id)
}
