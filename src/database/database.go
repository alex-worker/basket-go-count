package database

import (
	"basket-go-count/src/database/requests"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	db *sql.DB
}

type IDatabase interface {
	Connect(uri string) error
	Close()
	CalcSeasonRecords() (requests.SeasonRecords, error)
}

func New() IDatabase {
	return &Database{
		db: nil,
	}
}

func (i *Database) Connect(uri string) error {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return err
	}
	i.db = db
	return nil
}

func (i *Database) Close() {
	log.Print("Close db...")
	if i.db == nil {
		return
	}
	err := i.db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (i *Database) CalcSeasonRecords() (requests.SeasonRecords, error) {
	if i.db == nil {
		return nil, errors.New("not connected")
	}
	return requests.DoSeasonRecordsCount(i.db)
}
