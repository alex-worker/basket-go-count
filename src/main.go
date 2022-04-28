package main

import (
	"basket-go-count/src/config"
	"basket-go-count/src/informator"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	log.Println("Hello world! Lets go calculating...")
	config.Init()

	connStr := config.ReadConnectionString()
	log.Println("URI:", connStr)

	myInformator := informator.New()

	err := myInformator.Connect(connStr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer myInformator.Close()

	log.Println("Connect ok")

	records, err := myInformator.CalcSeasonRecords()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(records)
}
