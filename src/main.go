package main

import (
	"basket-go-count/src/config"
	"basket-go-count/src/database"
	"log"
)

func main() {
	log.Println("Hello world!")

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

	log.Println("+---------+--------+---------+")
	log.Println("| Season  | isUsa  | Players |")
	log.Println("+---------+--------+---------+")

	for _, value := range records {
		log.Printf("| %7v | %6v | %7v |", value.Season, value.IsUSA, value.Count)
	}

	log.Println("Good bye, cruel world")
}
