package db

import "github.com/jackc/pgx"

func NewPostgresDB(cfg pgx.ConnConfig) (*pgx.Conn, error) {
	db, err := pgx.Connect(cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}