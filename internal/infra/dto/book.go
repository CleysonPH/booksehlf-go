package dto

import (
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
