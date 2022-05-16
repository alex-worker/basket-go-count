package main

import (
	"basket-go-count/src/config"
	"basket-go-count/src/database"
	"log"
)

func main() {

	config.Init()
	connString := config.ReadConnectionString()

	myDb := database.New()

	err := myDb.Connect(connString)
	if err != nil {
		panic(err)
	}

	defer myDb.Close()

	records, err := myDb.CalcSeasonRecords()
	if err != nil {
		return
	}

	log.Println(records)

	log.Println("Hello world")
}
