package main

import (
	"basket-go-count/src/config"
	"basket-go-count/src/context"
	"basket-go-count/src/requests"
	"log"
)

func main() {
	log.Println("Hello world!")

	config.Init()
	connString := config.ReadConnectionString()

	ctx, err := context.NewContext(connString)

	if err != nil {
		panic(err)
	}

	defer ctx.Close()

	records, err := requests.DoSeasonRecordsCount(ctx.GetDb())
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
