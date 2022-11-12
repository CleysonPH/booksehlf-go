package dto

import (
	"testing"

	"github.com/cleysonph/bookshelf-go/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestBookSummaryResponse_FromDomain(t *testing.T) {
	book, _ := domain.NewBook(1, "Title", "9783161484100", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	wanted := &BookSummaryResponse{
		ID:         1,
		Title:      "Title",
		Isbn:       "9783161484100",
		Authors:    []string{"Author"},
		Categories: []string{"Category"},
		Language:   "en",
		Cover:      "http://example.com/cover.jpg",
	}
	got := &BookSummaryResponse{}
	got.FromDomain(book)
	assert.Equal(t, wanted, got)
}

func TestBookSummaryResponse_FromDomain_WithEmptyBook(t *testing.T) {
	book := &domain.Book{}
	wanted := &BookSummaryResponse{}
	got := &BookSummaryResponse{}
	got.FromDomain(book)
	assert.Equal(t, wanted, got)
}

func TestBookSummaryResponse_FromDomain_WithNilBook(t *testing.T) {
	wanted := &BookSummaryResponse{}
	got := &BookSummaryResponse{}
	got.FromDomain(nil)
	assert.Equal(t, wanted, got)
}
