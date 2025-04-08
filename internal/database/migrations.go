package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const folder = "internal/database/migrations"

func getMigrate() (*migrate.Migrate, error) {
	dsn := GetDSN()
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("failed to open database: ", err)
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("failed to create postgres driver instance: ", err)
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", folder),
		"postgres", driver,
	)

	if err != nil {
		log.Fatal("failed to create migrate instance: ", err)
		return nil, err
	}

	return m, nil
}

func Up() error {
	m, err := getMigrate()

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer m.Close()

	if err := m.Up(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func Down() error {
	m, err := getMigrate()

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer m.Close()

	if err := m.Down(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func Drop() error {
	m, err := getMigrate()

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer m.Close()

	if err := m.Drop(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func Reset() error {
	if err := Drop(); err != nil {
		log.Fatal(err)
		return err
	}

	if err := Up(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func NewMigration(name string) error {
	timestamp := time.Now().Format("20060102150405")
	upFile := fmt.Sprintf("%s/%s_%s.up.sql", folder, timestamp, name)
	downFile := fmt.Sprintf("%s/%s_%s.down.sql", folder, timestamp, name)

	if _, err := os.Create(upFile); err != nil {
		log.Fatal(err)
		return err
	}

	if _, err := os.Create(downFile); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
