package gateway

import "github.com/cleysonph/bookshelf-go/internal/domain"

type BookGateway interface {
	FindAllByTitle(title string) ([]*domain.Book, error)
}
