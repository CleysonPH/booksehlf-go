package validator

import (
	"fmt"
	"time"

	"github.com/cleysonph/bookshelf-go/internal/application/validator"
	"github.com/cleysonph/bookshelf-go/internal/domain"
)

type createBookValidator struct{}

// Validate implements validator.CreateBookValidator
func (*createBookValidator) Validate(book *domain.Book) error {
	err := validator.NewValidationError()

	err.AddErrorIf(book.Title() == "", "title", "Title is required")
	err.AddErrorIf(len(book.Title()) < 3, "title", "Title must be at least 3 characters long")
	err.AddErrorIf(len(book.Title()) > 255, "title", "Title must be at most 255 characters long")

	err.AddErrorIf(book.ISBN() == "", "isbn", "ISBN is required")
	err.AddErrorIf(len(book.ISBN()) != 13, "isbn", "ISBN must be 13 characters long")
	for _, c := range book.ISBN() {
		if c < '0' || c > '9' {
			err.AddError("isbn", "ISBN must contain only numbers")
			break
		}
	}

	err.AddErrorIf(len(book.Authors()) == 0, "authors", "Authors are required")
	for i, author := range book.Authors() {
		err.AddErrorIf(author == "", "author["+fmt.Sprintf("%d", i)+"]", "Author is required")
		err.AddErrorIf(len(author) < 3, "author["+fmt.Sprintf("%d", i)+"]", "Author must be at least 3 characters long")
	}

	err.AddErrorIf(len(book.Categories()) == 0, "categories", "Categories are required")
	for i, category := range book.Categories() {
		err.AddErrorIf(category == "", "category["+fmt.Sprintf("%d", i)+"]", "Category is required")
		err.AddErrorIf(len(category) < 3, "category["+fmt.Sprintf("%d", i)+"]", "Category must be at least 3 characters long")
	}

	err.AddErrorIf(book.Language() == "", "language", "Language is required")
	err.AddErrorIf(len(book.Language()) < 2, "language", "Language must be at least 2 characters long")

	if book.Description() != "" {
		err.AddErrorIf(len(book.Description()) < 10, "description", "Description must be at least 10 characters long")
		err.AddErrorIf(len(book.Description()) > 1000, "description", "Description must be at most 1000 characters long")
	}

	if !book.PublishedAt().IsZero() {
		err.AddErrorIf(book.PublishedAt().After(time.Now()), "publishedAt", "PublishedAt must be in the past")
	}

	if book.Publisher() != "" {
		err.AddErrorIf(len(book.Publisher()) < 3, "publisher", "Publisher must be at least 3 characters long")
		err.AddErrorIf(len(book.Publisher()) > 255, "publisher", "Publisher must be at most 255 characters long")
	}

	if book.Pages() != 0 {
		err.AddErrorIf(book.Pages() < 1, "pages", "Pages must be at least 1")
	}

	if book.Edition() != 0 {
		err.AddErrorIf(book.Edition() < 1, "edition", "Edition must be at least 1")
	}

	if err.HasErrors() {
		return err
	}
	return nil
}

func NewCreateBookValidator() validator.CreateBookValidator {
	return &createBookValidator{}
}
