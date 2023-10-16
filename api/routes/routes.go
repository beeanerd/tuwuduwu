package routes

import (
  "github.com/gorilla/mux"
  "github.com/beeanerd/tuwuduwu/handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

  r.HandleFunc("/healthcheck", handlers.HealthCheck).Methods("GET")

  r.HandleFunc("/addtodo", handlers.AddTodo).Methods("POST")

	return r
}
