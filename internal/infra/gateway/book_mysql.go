package gateway

import (
	"database/sql"
	"fmt"
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

func (b *bookTable) FromBook(book *domain.Book) {
	b.ID = book.ID()
	b.Title = book.Title()
	b.Isbn = book.ISBN()
	b.Authors = strings.Join(book.Authors(), ",")
	b.Categories = strings.Join(book.Categories(), ",")
	b.Language = book.Language()
	b.Cover = sql.NullString{String: book.Cover(), Valid: book.Cover() != ""}
	b.Description = sql.NullString{String: book.Description(), Valid: book.Description() != ""}
	b.PublishedAt = sql.NullTime{Time: book.PublishedAt(), Valid: !book.PublishedAt().IsZero()}
	b.Publisher = sql.NullString{String: book.Publisher(), Valid: book.Publisher() != ""}
	b.Pages = sql.NullInt32{Int32: book.Pages(), Valid: book.Pages() != 0}
	b.Edition = sql.NullInt32{Int32: book.Edition(), Valid: book.Edition() != 0}
	b.CreatedAt = book.CreatedAt()
	b.UpdatedAt = book.UpdatedAt()
}

func (b *bookTable) ToBook() (*domain.Book, error) {
	book, err := domain.NewBookWithAllValues(
		b.ID,
		b.Title,
		b.Isbn,
		strings.Split(b.Authors, ","),
		strings.Split(b.Categories, ","),
		b.Language,
		b.Cover.String,
		b.Description.String,
		b.PublishedAt.Time,
		b.Publisher.String,
		b.Pages.Int32,
		b.Edition.Int32,
		b.CreatedAt,
		b.UpdatedAt,
	)
	if err != nil {
		return nil, application.NewApplicationError(err, "error creating book: "+err.Error())
	}
	return book, nil
}

type BookMySQLGateway struct {
	db *sql.DB
}

const createQuery = `
INSERT INTO books (
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
	edition
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

// Create implements gateway.BookGateway
func (b *BookMySQLGateway) Create(book *domain.Book) (*domain.Book, error) {
	var f bookTable
	f.FromBook(book)
	result, err := db.Exec(
		createQuery,
		f.Title,
		f.Isbn,
		f.Authors,
		f.Categories,
		f.Language,
		f.Cover,
		f.Description,
		f.PublishedAt,
		f.Publisher,
		f.Pages,
		f.Edition,
	)
	if err != nil {
		return nil, application.NewApplicationError(err, "error creating book")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, application.NewApplicationError(err, "error getting last insert id")
	}
	return b.FindById(fmt.Sprintf("%d", id))
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
	return f.ToBook()
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
		book, err := f.ToBook()
		if err != nil {
			return nil, err
		}
		result = append(result, book)
	}
	return result, nil
}

func NewBookMySQLGateway(db *sql.DB) gateway.BookGateway {
	return &BookMySQLGateway{db: db}
}
