package validator

import (
	"fmt"
	"regexp"
	"time"

	"github.com/cleysonph/bookshelf-go/internal/application/validator"
	"github.com/cleysonph/bookshelf-go/internal/infra/dto"
)

type CreateBookValidator struct{}

func (*CreateBookValidator) Validate(book *dto.CreateBookRequest) error {
	err := validator.NewValidationError()

	err.AddErrorIf(book.Title == "", "title", "is required")
	err.AddErrorIf(len(book.Title) < 3, "title", "must be at least 3 characters long")
	err.AddErrorIf(len(book.Title) > 255, "title", "must be at most 255 characters long")

	// TODO: validate unique ISBN
	err.AddErrorIf(book.Isbn == "", "isbn", "is required")
	err.AddErrorIf(len(book.Isbn) != 13, "isbn", "must be 13 characters long")
	for _, c := range book.Isbn {
		if c < '0' || c > '9' {
			err.AddError("isbn", "must contain only numbers")
			break
		}
	}

	err.AddErrorIf(len(book.Authors) == 0, "authors", "are required")
	for i, author := range book.Authors {
		err.AddErrorIf(author == "", "author["+fmt.Sprintf("%d", i)+"]", "is required")
		err.AddErrorIf(len(author) < 3, "author["+fmt.Sprintf("%d", i)+"]", "must be at least 3 characters long")
	}

	err.AddErrorIf(len(book.Categories) == 0, "categories", "are required")
	for i, category := range book.Categories {
		err.AddErrorIf(category == "", "category["+fmt.Sprintf("%d", i)+"]", "is required")
		err.AddErrorIf(len(category) < 3, "category["+fmt.Sprintf("%d", i)+"]", "must be at least 3 characters long")
	}

	err.AddErrorIf(book.Language == "", "language", "is required")
	err.AddErrorIf(len(book.Language) != 2, "language", "must be 2 characters long")

	if book.Description.Value != "" {
		err.AddErrorIf(len(book.Description.Value) < 10, "description", "must be at least 10 characters long")
		err.AddErrorIf(len(book.Description.Value) > 1000, "description", "must be at most 1000 characters long")
	}

	if book.PublishedAt.Value != "" {
		isDateFmt, _ := regexp.MatchString(`^\d{4}-\d{2}-\d{2}$`, book.PublishedAt.Value)
		err.AddErrorIf(!isDateFmt, "published_at", "must be in the format YYYY-MM-DD")
		date, _ := time.Parse("2006-01-02", book.PublishedAt.Value)
		err.AddErrorIf(date.After(time.Now()), "published_at", "must be in the past")
	}

	if book.Publisher.Value != "" {
		err.AddErrorIf(len(book.Publisher.Value) < 3, "publisher", "must be at least 3 characters long")
		err.AddErrorIf(len(book.Publisher.Value) > 255, "publisher", "must be at most 255 characters long")
	}

	if book.Pages.Value != 0 {
		err.AddErrorIf(book.Pages.Value < 1, "pages", "must be at least 1")
	}

	if book.Edition.Value != 0 {
		err.AddErrorIf(book.Edition.Value < 1, "edition", "must be at least 1")
	}

	if err.HasErrors() {
		return err
	}
	return nil
}

func NewCreateBookValidator() validator.Validator[*dto.CreateBookRequest] {
	return &CreateBookValidator{}
}
