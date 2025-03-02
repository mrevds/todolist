package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// Load the .env file
	if err := godotenv.Load("internal/config/config.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create the data source name (DSN)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Open the database connection
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
	}

	// Ping the database to ensure the connection is established
	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Connected to the database successfully")
}

func GetDB() *sql.DB {
	return db
}