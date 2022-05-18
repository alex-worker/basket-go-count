package context

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Context struct {
	db *sql.DB
}

func NewContext(uri string) (*Context, error) {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Context{
		db,
	}, nil
}

func (ctx *Context) GetDb() *sql.DB {
	return ctx.db
}

func (ctx *Context) Close() {
	log.Print("Close db...")
	if ctx.db == nil {
		return
	}
	err := ctx.db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
