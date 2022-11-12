package gateway

import (
	"errors"
	"strings"

	"github.com/cleysonph/bookshelf-go/internal/application"
	"github.com/cleysonph/bookshelf-go/internal/application/gateway"
	"github.com/cleysonph/bookshelf-go/internal/domain"
)

type BookInMemoryGateway struct{}

// FindById implements gateway.BookGateway
func (*BookInMemoryGateway) FindById(id int64) (*domain.Book, error) {
	for _, book := range books {
		if book.ID() == id {
			return book, nil
		}
	}
	return nil, application.NewBookNotFoundError(errors.New("book not found"), "Book not found")
}

func initialBooks() []*domain.Book {
	books := make([]*domain.Book, 1)
	books[0], _ = domain.NewBook(1, "Title", "9783161484100", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	return books
}

var books = initialBooks()

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
