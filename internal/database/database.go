package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() (*sqlx.DB, error) {
	environment := os.Getenv("ENVIRONMENT")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	sslMode := ""
	if environment != "production" {
		sslMode = " sslmode=disable"
	}

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s%s",
		host,
		port,
		dbName,
		user,
		password,
		sslMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
		return nil, err
	}

	log.Println("Successfully Connected")

	return db, nil
}
