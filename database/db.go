package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)


func Connect() *sql.DB {
	connection := "user=Apolo dbname=goApi password=12345 host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}