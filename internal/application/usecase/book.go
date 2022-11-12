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
