package informator

import (
	"database/sql"
	"log"
)

const QuerySQL = `
	SELECT b.season, (CASE WHEN b.country = 'USA' THEN 'USA' ELSE 'Others' END) as is_usa, count(b.player_name)
	FROM basketball b
	GROUP BY b.season, is_usa
	ORDER BY season, is_usa DESC;
`

type SeasonRecord struct {
	Season string
	IsUSA  string
	Count  int
}

type SeasonRecords []SeasonRecord

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
	rows, err := i.db.Query(QuerySQL)
	if err != nil {
		return nil, err
	}

	var resp SeasonRecords

	for rows.Next() {
		var season string
		var isUSA string
		var count int

		err := rows.Scan(&season, &isUSA, &count)
		switch err {
		case sql.ErrNoRows:
			return resp, nil
		case nil:
			resp = append(resp, SeasonRecord{
				season,
				isUSA,
				count,
			})
		default:
			return resp, err
		}
	}

	err = rows.Close()
	if err != nil {
		return resp, err
	}
	return resp, nil
}
