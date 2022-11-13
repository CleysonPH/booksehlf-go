package gateway

import "github.com/cleysonph/bookshelf-go/internal/domain"

type BookGateway interface {
	FindAllByTitle(title string) ([]*domain.Book, error)
	FindById(id string) (*domain.Book, error)
	DeleteById(id string) error
	ExistsById(id string) bool
}
