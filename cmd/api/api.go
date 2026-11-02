package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucaastorres7/ecom-api/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

// New function for the API Server
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// API Server Methods

/*
1. Init Router
2. Register Routers
3. Register Dependencies
*/
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	// Services
	userHandler := user.NewHandler(
		userStore,
	)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
