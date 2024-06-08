package api

import (
	"database/sql"
	"net/http"
	"project/restful-api/services"

	"github.com/gorilla/mux"
)

type API struct {
	addr string
	db   *sql.DB
}

func New(addr string, db *sql.DB) *API {
	return &API{
		addr: addr,
		db:   db,
	}
}

func (a *API) Start() error {
	// Start the server
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Register the handlers
	userHandler := services.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	return http.ListenAndServe(a.addr, router)
}