package validator

import "github.com/cleysonph/bookshelf-go/internal/domain"

type CreateBookValidator interface {
	Validate(book *domain.Book) error
}
