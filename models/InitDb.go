package models

import (
	r "github.com/dancannon/gorethink"
	"log"
	"time"
)

var Session *r.Session

func InitDb() {
	//Setup RethinkDB connection pool
	session, err := r.Connect(r.ConnectOpts{
		Address:     "localhost:28015",
		Database:    "choreboard",
		MaxIdle:     10,
		IdleTimeout: time.Second * 10,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	Session = session
}
