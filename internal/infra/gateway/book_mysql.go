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
	Cover       string
	Description string
	PublishedAt time.Time
	Publisher   string
	Pages       int32
	Edition     int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type BookMySQLGateway struct {
	db *sql.DB
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
func (*BookMySQLGateway) FindById(id int64) (*domain.Book, error) {
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
		return nil, application.NewApplicationError(err, "error scanning database row")
	}
	book, err := domain.NewBookWithAllValues(
		f.ID,
		f.Title,
		f.Isbn,
		strings.Split(f.Authors, ","),
		strings.Split(f.Categories, ","),
		f.Language,
		f.Cover,
		f.Description,
		f.PublishedAt,
		f.Publisher,
		f.Pages,
		f.Edition,
		f.CreatedAt,
		f.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, application.NewBookNotFoundError(err, "Book not found")
		}
		return nil, application.NewApplicationError(err, "error creating book")
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
			f.Cover,
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
