package api

import "database/sql"

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