package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ecomerce-api/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	// Create a subrouter for /api/v1
	apiRouter := http.NewServeMux()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(apiRouter)

	// Add the subrouter to the main router with the prefix
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRouter))

	log.Println("Starting server on", s.addr)
	return http.ListenAndServe(s.addr, nil)
}
