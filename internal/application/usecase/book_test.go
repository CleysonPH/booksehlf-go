package usecase

import (
	"testing"

	"github.com/cleysonph/bookshelf-go/internal/application"
	"github.com/cleysonph/bookshelf-go/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type bookGatewayMock struct {
	mock.Mock
}

// Create implements gateway.BookGateway
func (b *bookGatewayMock) Create(book *domain.Book) (*domain.Book, error) {
	args := b.Called(book)
	return args.Get(0).(*domain.Book), args.Error(1)
}

// ExistsById implements gateway.BookGateway
func (b *bookGatewayMock) ExistsById(id string) bool {
	args := b.Called(id)
	return args.Bool(0)
}

// DeleteById implements gateway.BookGateway
func (b *bookGatewayMock) DeleteById(id string) error {
	args := b.Called(id)
	return args.Error(0)
}

// FindById implements gateway.BookGateway
func (b *bookGatewayMock) FindById(id string) (*domain.Book, error) {
	args := b.Called(id)
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (b *bookGatewayMock) FindAllByTitle(title string) ([]*domain.Book, error) {
	args := b.Called(title)
	return args.Get(0).([]*domain.Book), args.Error(1)
}

func TestGetBooksUseCase_Execute_ShouldReturnBookList(t *testing.T) {
	bookGateway := &bookGatewayMock{}
	expectedBook, _ := domain.NewBook(1, "Title", "9783161484100", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	bookGateway.On("FindAllByTitle", "Title").Return([]*domain.Book{expectedBook}, nil)
	usecase := NewGetBooksUseCase(bookGateway)
	books, err := usecase.Execute("Title")
	assert.Nil(t, err)
	assert.Equal(t, []*domain.Book{expectedBook}, books)
}

func TestGetBooksUseCase_Execute_ShouldReturnError(t *testing.T) {
	bookGateway := &bookGatewayMock{}
	bookGateway.On("FindAllByTitle", "Title").Return([]*domain.Book{}, application.NewApplicationError(nil, "Error"))
	usecase := NewGetBooksUseCase(bookGateway)
	books, err := usecase.Execute("Title")
	assert.Equal(t, application.NewApplicationError(nil, "Error"), err)
	assert.Empty(t, books)
}
