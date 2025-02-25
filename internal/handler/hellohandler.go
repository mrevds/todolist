package handler

import (
	"net/http"
	"encoding/json"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello, API is working!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}