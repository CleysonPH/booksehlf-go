package usecase

import (
	"testing"

	"github.com/cleysonph/bookshelf-go/internal/application"
	"github.com/cleysonph/bookshelf-go/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BookGatewayMock struct {
	mock.Mock
}

func (m *BookGatewayMock) FindAllByTitle(title string) ([]*domain.Book, error) {
	args := m.Called(title)
	return args.Get(0).([]*domain.Book), args.Error(1)
}

func TestGetBooksUseCase_Execute_ShouldReturnBookList(t *testing.T) {
	bookGateway := &BookGatewayMock{}
	expectedBook, _ := domain.NewBook(1, "Title", "9783161484100", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	bookGateway.On("FindAllByTitle", "Title").Return([]*domain.Book{expectedBook}, nil)
	usecase := NewGetBooksUseCase(bookGateway)
	books, err := usecase.Execute("Title")
	assert.Nil(t, err)
	assert.Equal(t, []*domain.Book{expectedBook}, books)
}

func TestGetBooksUseCase_Execute_ShouldReturnError(t *testing.T) {
	bookGateway := &BookGatewayMock{}
	bookGateway.On("FindAllByTitle", "Title").Return([]*domain.Book{}, application.NewApplicationError(nil, "Error"))
	usecase := NewGetBooksUseCase(bookGateway)
	books, err := usecase.Execute("Title")
	assert.Equal(t, application.NewApplicationError(nil, "Error"), err)
	assert.Empty(t, books)
}