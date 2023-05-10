package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

const (
	devDB = "postgresql://postgres:password@localhost:5432/go-chat?sslmode=disable"
	prodDB = "postgresql://root:password@localhost:5433/go-chat?sslmode=disable"
)

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres", prodDB)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}