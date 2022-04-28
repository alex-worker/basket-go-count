package informator

import (
	"database/sql"
	"log"
)

func CloseDb(db *sql.DB) {
	log.Print("Close db...")
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func CloseRows(rows *sql.Rows) {
	log.Print("Close rows...")
	err := rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}
