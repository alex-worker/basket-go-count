package main

import (
	"basket-go-count/src/config"
	"basket-go-count/src/database"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	log.Println("Hello world! Lets go calculating...")
	config.Init()

	connStr := config.ReadConnectionString()
	log.Println("URI:", connStr)

	myDatabase := database.New()

	err := myDatabase.Connect(connStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer myDatabase.Close()

	log.Println("Connect ok")

	records, err := myDatabase.CalcSeasonRecords()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		log.Printf(" %v %v %v", record.Season, record.IsUSA, record.Count)
	}

	log.Println("Bye-bue, cruel world!...")
}
