package main

import (
	"log"
	"net/http"
	"todo/internal/router"
)

func main() {
	r := router.NewRouter()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}