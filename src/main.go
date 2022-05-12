package main

import (
	"database/sql"
	_ "github.com/lib/pq" // add this
	"log"
)

func main() {
	log.Println("Hello world")

	connStr := "postgresql://postgres:postgres@localhost:5432/todos?sslmode=disable"
	// Connect to database
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Ok")
}
