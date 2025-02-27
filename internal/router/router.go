package router

import (
	"github.com/gorilla/mux"
	"todo/internal/handler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/hello", handler.HelloHandler).Methods("GET")
	r.HandleFunc("/v1/task", handler.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/v1/task/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	return r
}
