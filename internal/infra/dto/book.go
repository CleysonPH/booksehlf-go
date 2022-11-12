package dto

import (
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
