package main

import (
	"basket-go-count/src/config"
	"basket-go-count/src/informator"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type record struct {
	season string
	isUSA  string
	count  int
}

func main() {

	config.Read()

	//log.Println("Hello world! Lets go calculating...")
	//log.Println("Connect..")

	connStr := config.ReadConnectionString()

	log.Println("URI:", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer informator.CloseDb(db)

	log.Println("Connect ok")

	rows, err := db.Query(informator.QuerySQL)

	log.Printf("%T", err)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		p := record{}
		err := rows.Scan(&p.season, &p.isUSA, &p.count)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("Empty result!")
			return
		case nil:
			fmt.Printf("%v %v %v\n", p.season, p.isUSA, p.count)
		default:
			log.Fatal(err)
		}
	}

	informator.CloseRows(rows)
}
