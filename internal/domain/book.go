package domain

import (
	"errors"
	"time"
)

type Book struct {
	id          int64
	title       string
	isbn        string
	authors     []string
	categories  []string
	language    string
	cover       string
	description string
	publishedAt time.Time
	publisher   string
	pages       int32
	edition     int32
	createdAt   time.Time
	updatedAt   time.Time
}

func NewBook(
	id int64,
	title string,
	isbn string,
	authors []string,
	categories []string,
	language string,
	cover string,
) (*Book, error) {
	b := &Book{
		id:         id,
		title:      title,
		isbn:       isbn,
		authors:    authors,
		categories: categories,
		language:   language,
		cover:      cover,
	}
	if err := b.validate(); err != nil {
		return nil, err
	}
	return b, nil
}

func NewBookWithAllValues(
	id int64,
	title string,
	isbn string,
	authors []string,
	categories []string,
	language string,
	cover string,
	description string,
	publishedAt time.Time,
	publisher string,
	pages int32,
	edition int32,
	createdAt time.Time,
	updatedAt time.Time,
) (*Book, error) {
	b := &Book{
		id:          id,
		title:       title,
		isbn:        isbn,
		authors:     authors,
		categories:  categories,
		language:    language,
		cover:       cover,
		description: description,
		publishedAt: publishedAt,
		publisher:   publisher,
		pages:       pages,
		edition:     edition,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
	if err := b.validate(); err != nil {
		return nil, err
	}
	return b, nil
}

func (b Book) ID() int64 {
	return b.id
}

func (b Book) Title() string {
	return b.title
}

func (b Book) ISBN() string {
	return b.isbn
}

func (b Book) Authors() []string {
	return b.authors
}

func (b Book) Categories() []string {
	return b.categories
}

func (b Book) Language() string {
	return b.language
}

func (b Book) Cover() string {
	return b.cover
}

func (b Book) Description() string {
	return b.description
}

func (b Book) PublishedAt() time.Time {
	return b.publishedAt
}

func (b Book) Publisher() string {
	return b.publisher
}

func (b Book) Pages() int32 {
	return b.pages
}

func (b Book) Edition() int32 {
	return b.edition
}

func (b Book) CreatedAt() time.Time {
	return b.createdAt
}

func (b Book) UpdatedAt() time.Time {
	return b.updatedAt
}

func (b Book) validate() error {
	if b.title == "" {
		return errors.New("title is required")
	}

	if b.isbn == "" {
		return errors.New("isbn is required")
	}

	if len(b.isbn) != 13 {
		return errors.New("isbn must be 13 digits")
	}

	for _, c := range b.isbn {
		if c < '0' || c > '9' {
			return errors.New("isbn must contain only numbers")
		}
	}

	if len(b.authors) == 0 {
		return errors.New("authors are required")
	}

	if len(b.categories) == 0 {
		return errors.New("categories are required")
	}

	if b.language == "" {
		return errors.New("language is required")
	}

	if len(b.language) != 2 {
		return errors.New("language must be 2 characters")
	}

	return nil
}
