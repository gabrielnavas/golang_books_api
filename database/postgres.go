package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func MakePostgresSQLDatabase() *sql.DB {
	connStr := "host=localhost user=postgres dbname=book_api " +
		"password=postgres123 port=5433 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[ * ] DATABASE IS ON")
	return db
}
