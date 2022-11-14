package usecase

import (
	"errors"

	"github.com/cleysonph/bookshelf-go/internal/application"
	"github.com/cleysonph/bookshelf-go/internal/application/gateway"
	"github.com/cleysonph/bookshelf-go/internal/application/validator"
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
	if !u.bookGateway.ExistsById(id) {
		return application.NewBookNotFoundError(errors.New("book not found"), "Book not found")
	}
	return u.bookGateway.DeleteById(id)
}

func NewCreateBookUseCase(
	bookGateway gateway.BookGateway,
	createBookValidator validator.CreateBookValidator,
) *CreateBookUseCase {
	return &CreateBookUseCase{bookGateway, createBookValidator}
}

type CreateBookUseCase struct {
	bookGateway         gateway.BookGateway
	createBookValidator validator.CreateBookValidator
}

func (u *CreateBookUseCase) Execute(createBook *domain.Book) (*domain.Book, error) {
	if err := u.createBookValidator.Validate(createBook); err != nil {
		return nil, err
	}
	return u.bookGateway.Create(createBook)
}
