package informator

import (
	"database/sql"
	"errors"
	"log"
)

type Informator struct {
	db *sql.DB
}

type IInformator interface {
	Connect(uri string) error
	Close()
	CalcSeasonRecords() (SeasonRecords, error)
}

func New() IInformator {
	return &Informator{
		db: nil,
	}
}

func (i *Informator) Connect(uri string) error {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return err
	}
	i.db = db
	return nil
}

func (i *Informator) Close() {
	log.Print("Close db...")
	if i.db == nil {
		return
	}
	err := i.db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func (i *Informator) CalcSeasonRecords() (SeasonRecords, error) {
	if i.db == nil {
		return nil, errors.New("not connected")
	}
	return DoSeasonRecordsCount(i.db)
}
