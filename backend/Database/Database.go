package db

import (
	"database/sql"
	"log"
	// "github.com/lib/pq"
)

func main() {
	connStr := "postgres://postgres:JessePostgres2003?!@192.168.0.11:5432/MovieApp?sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	log.Print(rows)
}

func Query() *sql.Rows {
	connStr := "postgres://postgres:JessePostgres2003?!@192.168.0.11:5432/MovieApp?sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	age := 21
	rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
	return rows
}