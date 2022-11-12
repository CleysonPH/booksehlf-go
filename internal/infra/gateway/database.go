package gateway

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	err error
)

func InitDB(dsn string, dbdriver string) {
	db, err = sql.Open(dbdriver, dsn)
	if err != nil {
		log.Fatal(err)
	}
}

func DB() *sql.DB {
	return db
}

func CloseDB() {
	db.Close()
}
