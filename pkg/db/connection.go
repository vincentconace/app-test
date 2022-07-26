package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "../app-test.db")
	if err != nil {
		panic(err)
	}

	return db
}
