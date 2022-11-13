package gateway

import (
	"database/sql"
	"strings"
	"time"

	"github.com/cleysonph/bookshelf-go/internal/application"
	"github.com/cleysonph/bookshelf-go/internal/application/gateway"
	"github.com/cleysonph/bookshelf-go/internal/domain"
)

type bookTable struct {
	ID          int64
	Title       string
	Isbn        string
	Authors     string
	Categories  string
	Language    string
	Cover       sql.NullString
	Description sql.NullString
	PublishedAt sql.NullTime
	Publisher   sql.NullString
	Pages       sql.NullInt32
	Edition     sql.NullInt32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookMySQLGateway struct {
	db *sql.DB
}

const existsByIdQuery = `
SELECT
	id
FROM
	books
WHERE
	id = ?
`

// ExistsById implements gateway.BookGateway
func (*BookMySQLGateway) ExistsById(id string) bool {
	row := db.QueryRow(existsByIdQuery, id)
	var f int64
	if err := row.Scan(&f); err != nil {
		return false
	}
	return true
}

const deleteByIdQuery = `
DELETE FROM
	books
WHERE
	id = ?
`

// DeleteById implements gateway.BookGateway
func (*BookMySQLGateway) DeleteById(id string) error {
	if _, err := db.Exec(deleteByIdQuery, id); err != nil {
		return application.NewApplicationError(err, "error deleting book")
	}
	return nil
}

const findByIdQuery = `
SELECT
	id,
	title,
	isbn,
	authors,
	categories,
	language,
	cover,
	description,
	published_at,
	publisher,
	pages,
	edition,
	created_at,
	updated_at
FROM
	books
WHERE
	id = ?
LIMIT 1
`

// FindById implements gateway.BookGateway
func (*BookMySQLGateway) FindById(id string) (*domain.Book, error) {
	row := db.QueryRow(findByIdQuery, id)
	var f bookTable
	if err := row.Scan(
		&f.ID,
		&f.Title,
		&f.Isbn,
		&f.Authors,
		&f.Categories,
		&f.Language,
		&f.Cover,
		&f.Description,
		&f.PublishedAt,
		&f.Publisher,
		&f.Pages,
		&f.Edition,
		&f.CreatedAt,
		&f.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, application.NewBookNotFoundError(err, "Book not found")
		}
		return nil, application.NewApplicationError(err, "error scanning database row")
	}
	book, err := domain.NewBookWithAllValues(
		f.ID,
		f.Title,
		f.Isbn,
		strings.Split(f.Authors, ","),
		strings.Split(f.Categories, ","),
		f.Language,
		f.Cover.String,
		f.Description.String,
		f.PublishedAt.Time,
		f.Publisher.String,
		f.Pages.Int32,
		f.Edition.Int32,
		f.CreatedAt,
		f.UpdatedAt,
	)
	if err != nil {
		return nil, application.NewApplicationError(err, "error creating book: "+err.Error())
	}
	return book, nil
}

const findAllByTitleQuery = `
SELECT
	id,
	title,
	isbn,
	authors,
	categories,
	language,
	cover
FROM
	books
WHERE
	LOWER(title) LIKE CONCAT('%', LOWER(?), '%')
`

// FindAllByTitle implements gateway.BookGateway
func (*BookMySQLGateway) FindAllByTitle(title string) ([]*domain.Book, error) {
	rows, err := db.Query(findAllByTitleQuery, title)
	if err != nil {
		return nil, application.NewApplicationError(err, "error querying database")
	}
	defer rows.Close()

	var result []*domain.Book
	for rows.Next() {
		var f bookTable
		if err := rows.Scan(
			&f.ID,
			&f.Title,
			&f.Isbn,
			&f.Authors,
			&f.Categories,
			&f.Language,
			&f.Cover,
		); err != nil {
			return nil, application.NewApplicationError(err, "error scanning database row")
		}
		book, err := domain.NewBook(
			f.ID,
			f.Title,
			f.Isbn,
			strings.Split(f.Authors, ","),
			strings.Split(f.Categories, ","),
			f.Language,
			f.Cover.String,
		)
		if err != nil {
			return nil, application.NewApplicationError(err, "error creating book")
		}
		result = append(result, book)
	}
	return result, nil
}

func NewBookMySQLGateway(db *sql.DB) gateway.BookGateway {
	return &BookMySQLGateway{db: db}
}
