package gateway

import (
	"database/sql"
	"strings"

	"github.com/cleysonph/bookshelf-go/internal/application"
	"github.com/cleysonph/bookshelf-go/internal/application/gateway"
	"github.com/cleysonph/bookshelf-go/internal/domain"
)

type BookMySQLGateway struct {
	db *sql.DB
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
	type fields struct {
		ID         int64
		Title      string
		Isbn       string
		Authors    string
		Categories string
		Language   string
		Cover      string
	}
	rows, err := db.Query(findAllByTitleQuery, title)
	if err != nil {
		return nil, application.NewApplicationError(err, "error querying database")
	}
	defer rows.Close()

	var result []*domain.Book
	for rows.Next() {
		var f fields
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
