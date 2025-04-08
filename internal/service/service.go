package service

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func NewPool(dbPool *sqlx.DB) error {
	db = dbPool
	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}
