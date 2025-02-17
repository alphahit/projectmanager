package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store
}

// Constructore for the service
func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

// TO INITIALIZE THE ROUTER
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	//registering our services
	tasksService := NewTasksService(s.store)
	tasksService.RegisterRoutes(router)

	log.Println("Starting the API server at", s.addr)
	log.Fatal(http.ListenAndServe(s.addr, subrouter))
}
