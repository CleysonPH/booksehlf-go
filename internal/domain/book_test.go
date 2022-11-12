package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewBook_WithEmptyTitle_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "", "9783161484100", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "title is required")
}

func TestNewBook_WithEmptyISBN_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "Title", "", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "isbn is required")
}

func TestNewBook_WithLettersInISBN_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "Title", "978316148410a", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "isbn must contain only numbers")
}

func TestNewBook_WithSpacesInISBN_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "Title", "978 316148410", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "isbn must contain only numbers")
}

func TestNewBook_WithSpecialCharactersInISBN_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "Title", "978-316148410", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "isbn must contain only numbers")
}

func TestNewBook_WithMoreThan13DigitsInISBN_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "Title", "97831614841001", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "isbn must be 13 digits")
}

func TestNewBook_WithLessThan13DigitsInISBN_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "Title", "978316148410", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "isbn must be 13 digits")
}

func TestNewBook_WithEmptyAuthors_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "Title", "9783161484100", []string{}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "authors are required")
}

func TestNewBook_WithEmptyCategories_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "Title", "9783161484100", []string{"Author"}, []string{}, "en", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "categories are required")
}

func TestNewBook_WithEmptyLanguage_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "Title", "9783161484100", []string{"Author"}, []string{"Category"}, "", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "language is required")
}

func TestNewBook_WithMoreThan2CharactersLanguage_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "Title", "9783161484100", []string{"Author"}, []string{"Category"}, "eng", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "language must be 2 characters")
}

func TestNewBook_WithLessThan2CharactersLanguage_ShouldReturnError(t *testing.T) {
	_, err := NewBook(1, "Title", "9783161484100", []string{"Author"}, []string{"Category"}, "e", "http://example.com/cover.jpg")
	assert.ErrorContains(t, err, "language must be 2 characters")
}

func TestNewBook_WithValidData_ShouldReturnBook(t *testing.T) {
	book, err := NewBook(1, "Title", "9783161484100", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	assert.Nil(t, err)
	assert.Equal(t, book.ID(), int64(1))
	assert.Equal(t, book.Title(), "Title")
	assert.Equal(t, book.ISBN(), "9783161484100")
	assert.Equal(t, book.Authors(), []string{"Author"})
	assert.Equal(t, book.Categories(), []string{"Category"})
	assert.Equal(t, book.Language(), "en")
	assert.Equal(t, book.Cover(), "http://example.com/cover.jpg")
}

func TestNewBookWithAllValues_WithInvalidData_ShouldReturnError(t *testing.T) {
	_, err := NewBookWithAllValues(
		1,
		"",
		"9783161484100",
		[]string{"Author"},
		[]string{"Category"},
		"en",
		"http://example.com/cover.jpg",
		"Description",
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		"Publisher",
		10,
		1,
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	)
	assert.ErrorContains(t, err, "title is required")
}

func TestNewBookWithAllValues_WithValidData_ShouldReturnBook(t *testing.T) {
	book, err := NewBookWithAllValues(
		1,
		"Title",
		"9783161484100",
		[]string{"Author"},
		[]string{"Category"},
		"en",
		"http://example.com/cover.jpg",
		"Description",
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		"Publisher",
		10,
		1,
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	)
	assert.Nil(t, err)
	assert.Equal(t, book.ID(), int64(1))
	assert.Equal(t, book.Title(), "Title")
	assert.Equal(t, book.ISBN(), "9783161484100")
	assert.Equal(t, book.Authors(), []string{"Author"})
	assert.Equal(t, book.Categories(), []string{"Category"})
	assert.Equal(t, book.Language(), "en")
	assert.Equal(t, book.Cover(), "http://example.com/cover.jpg")
	assert.Equal(t, book.Description(), "Description")
	assert.Equal(t, book.PublishedAt(), time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, book.Publisher(), "Publisher")
	assert.Equal(t, book.Pages(), int32(10))
	assert.Equal(t, book.Edition(), int32(1))
	assert.Equal(t, book.CreatedAt(), time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
	assert.Equal(t, book.UpdatedAt(), time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC))
}
