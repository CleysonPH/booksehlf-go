package controller

import (
	"errors"
	"net/url"
	"strings"
	"testing"

	"github.com/cleysonph/bookshelf-go/internal/application"
	"github.com/cleysonph/bookshelf-go/internal/application/usecase"
	"github.com/cleysonph/bookshelf-go/internal/domain"
	"github.com/cleysonph/bookshelf-go/internal/infra/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type bookGatewayMock struct {
	mock.Mock
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

func (m *bookGatewayMock) FindAllByTitle(title string) ([]*domain.Book, error) {
	args := m.Called(title)
	return args.Get(0).([]*domain.Book), args.Error(1)
}

func TestGetBooksWebController_Execute(t *testing.T) {
	gateway := &bookGatewayMock{}
	usecase := usecase.NewGetBooksUseCase(gateway)
	controller := NewGetBooksWebController(usecase)

	expectedBook, _ := domain.NewBook(1, "Title", "9783161484100", []string{"Author"}, []string{"Category"}, "en", "http://example.com/cover.jpg")
	gateway.On("FindAllByTitle", "Title").Return([]*domain.Book{expectedBook}, nil)

	resp := controller.Execute(web.NewHttpRequest(url.Values{"q": []string{"Title"}}, map[string]string{}))
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Headers.ContentType)
	assert.True(t, strings.Contains(string(resp.Body), `"title":"Title"`))
	assert.True(t, strings.Contains(string(resp.Body), `"isbn":"9783161484100"`))
}

func TestGetBooksWebController_ExecuteWithError(t *testing.T) {
	gateway := &bookGatewayMock{}
	usecase := usecase.NewGetBooksUseCase(gateway)
	controller := NewGetBooksWebController(usecase)

	gateway.On("FindAllByTitle", "Title").Return([]*domain.Book{}, application.NewApplicationError(errors.New("error"), "Error"))

	resp := controller.Execute(web.NewHttpRequest(url.Values{"q": []string{"Title"}}, map[string]string{}))
	assert.Equal(t, 500, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Headers.ContentType)
	assert.True(t, strings.Contains(string(resp.Body), `"error":"Internal Server Error"`))
}
