package main

import (
	"log"
	"net/http"
	"todo/internal/database"
	"todo/internal/router"
)

func main() {
	_, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	r := router.NewRouter()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
