package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	dsn := "user=postgres password=12345678 dbname=todo sslmode=disable" // replace with your database credentials
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to the database successfully")
	return db, nil
}

func GetDB() *sql.DB {
	return db
}
