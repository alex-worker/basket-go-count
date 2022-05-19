package mocks

import (
	"database/sql"
	_ "github.com/lib/pq"
)

//func GetSQL_DB() *sql.DB {
//	db, err := sql.Open("postgres", GetConnectionString())
//	if err != nil {
//		panic(err)
//	}
//	return db
//}

func Connect() *sql.DB {
	db, err := sql.Open("postgres", GetConnectionStringRealDb())
	if err != nil {
		panic(err)
	}
	return db
}
