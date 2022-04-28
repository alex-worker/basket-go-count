package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

const deafaultConnectString = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

func Read() {
	var name string
	flag.StringVar(&name, "uri", deafaultConnectString, "Specify name. Default is admin.")

	flag.Usage = func() {
		fmt.Printf("Usage of basket-go-count: \n")
		fmt.Printf("./basket-go-count -uri <db connect uri>\n")
		fmt.Printf("<db connect uri> = <postgres>://<username>:<password>@<host>:<port>/<database>[?sslmode=<sslmode>]\n")
	}
	flag.Parse()

	log.Print(name)
}

func ReadConnectionString() string {
	log.Print(os.Getenv("uri"))
	return viper.GetString("db.Uri")
	//connStr := "host=balarama.db.elephantsql.com port=5432 user=mmxxboti password=18iR36iF2GJi1kY4ki7HfHw4Idamq1_M dbname=mmxxboti"
	//const connStr = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
}
