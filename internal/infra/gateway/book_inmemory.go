package gateway

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cleysonph/bookshelf-go/internal/application"
	"github.com/cleysonph/bookshelf-go/internal/application/gateway"
	"github.com/cleysonph/bookshelf-go/internal/domain"
)

func initialBooks() []*domain.Book {
	books := make([]*domain.Book, 1)
	books[0], _ = domain.NewBook(1, "Title", "9783161484100", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	return books
}

var books = initialBooks()

type BookInMemoryGateway struct{}

// ExistsById implements gateway.BookGateway
func (*BookInMemoryGateway) ExistsById(id string) bool {
	for _, book := range books {
		if fmt.Sprintf("%d", book.ID()) == id {
			return true
		}
	}
	return false
}

// DeleteById implements gateway.BookGateway
func (*BookInMemoryGateway) DeleteById(id string) error {
	for i, book := range books {
		if fmt.Sprintf("%d", book.ID()) == id {
			books = append(books[:i], books[i+1:]...)
			return nil
		}
	}
	return application.NewBookNotFoundError(errors.New("book not found"), "Book not found")
}

// FindById implements gateway.BookGateway
func (*BookInMemoryGateway) FindById(id string) (*domain.Book, error) {
	for _, book := range books {
		if fmt.Sprintf("%d", book.ID()) == id {
			return book, nil
		}
	}
	return nil, application.NewBookNotFoundError(errors.New("book not found"), "Book not found")
}

// FindAllByTitle implements gateway.BookGateway
func (*BookInMemoryGateway) FindAllByTitle(title string) ([]*domain.Book, error) {
	var result []*domain.Book
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title()), strings.ToLower(title)) {
			result = append(result, book)
		}
	}
	return result, nil
}

func NewBookInMemoryGateway() gateway.BookGateway {
	return &BookInMemoryGateway{}
}
