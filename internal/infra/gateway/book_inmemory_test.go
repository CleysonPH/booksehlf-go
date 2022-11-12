package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookInMemoryGateway_FindAllByTitle(t *testing.T) {
	gateway := NewBookInMemoryGateway()
	books, err := gateway.FindAllByTitle("Title")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(books))
}
