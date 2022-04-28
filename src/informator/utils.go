package informator

import (
	"database/sql"
	"log"
)

func closeRows(rows *sql.Rows) {
	log.Print("Close rows...")
	err := rows.Close()
	if err != nil {
		log.Fatal(err)
	}
}
