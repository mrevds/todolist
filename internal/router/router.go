package router

import (
	"github.com/gorilla/mux"
	"todo/internal/handler"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/v1/hello", handler.HelloHandler).Methods("GET")
	r.HandleFunc("/task", handler.CreateTaskHandler).Methods("POST")
	r.HandleFunc("/task/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/task/{id}", handler.UpdateTaskHandler).Methods("PUT")
	r.HandleFunc("/task/{id}", handler.GetTaskHandler).Methods("GET")
	r.HandleFunc("/tasks", handler.GetTasksHandler).Methods("GET")
	return r
}
