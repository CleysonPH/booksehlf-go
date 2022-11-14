package dto

import (
	"encoding/json"
	"time"

	"github.com/cleysonph/bookshelf-go/internal/domain"
)

type BookSummaryResponse struct {
	ID         int64    `json:"id"`
	Title      string   `json:"title"`
	Isbn       string   `json:"isbn"`
	Authors    []string `json:"authors"`
	Categories []string `json:"categories"`
	Language   string   `json:"language"`
	Cover      string   `json:"cover"`
}

func (b *BookSummaryResponse) FromDomain(book *domain.Book) {
	if book == nil {
		return
	}
	b.ID = book.ID()
	b.Title = book.Title()
	b.Isbn = book.ISBN()
	b.Authors = book.Authors()
	b.Categories = book.Categories()
	b.Language = book.Language()
	b.Cover = book.Cover()
}

type BookResponse struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Isbn        string     `json:"isbn"`
	Authors     []string   `json:"authors"`
	Categories  []string   `json:"categories"`
	Language    string     `json:"language"`
	Cover       NullString `json:"cover"`
	Description NullString `json:"description"`
	PublishedAt NullTime   `json:"published_at"`
	Publisher   NullString `json:"publisher"`
	Pages       NullInt32  `json:"pages"`
	Edition     NullInt32  `json:"edition"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (b *BookResponse) FromDomain(book *domain.Book) {
	if book == nil {
		return
	}
	b.ID = book.ID()
	b.Title = book.Title()
	b.Isbn = book.ISBN()
	b.Authors = book.Authors()
	b.Categories = book.Categories()
	b.Language = book.Language()
	b.Cover = NullString{Value: book.Cover()}
	b.Description = NullString{Value: book.Description()}
	b.PublishedAt = NullTime{Value: book.PublishedAt()}
	b.Publisher = NullString{Value: book.Publisher()}
	b.Pages = NullInt32{Value: book.Pages()}
	b.Edition = NullInt32{Value: book.Edition()}
	b.CreatedAt = book.CreatedAt()
	b.UpdatedAt = book.UpdatedAt()
}

type CreateBookRequest struct {
	Title       string     `json:"title"`
	Isbn        string     `json:"isbn"`
	Authors     []string   `json:"authors"`
	Categories  []string   `json:"categories"`
	Language    string     `json:"language"`
	Description NullString `json:"description"`
	PublishedAt NullTime   `json:"published_at"`
	Publisher   NullString `json:"publisher"`
	Pages       NullInt32  `json:"pages"`
	Edition     NullInt32  `json:"edition"`
}

func (b *CreateBookRequest) FromJson(jsonBody []byte) error {
	return json.Unmarshal(jsonBody, b)
}

func (b *CreateBookRequest) ToDomain() *domain.Book {
	book, _ := domain.NewBookWithAllValues(
		0,
		b.Title,
		b.Isbn,
		b.Authors,
		b.Categories,
		b.Language,
		"",
		b.Description.Value,
		b.PublishedAt.Value,
		b.Publisher.Value,
		b.Pages.Value,
		b.Edition.Value,
		time.Time{},
		time.Time{},
	)
	return book
}
