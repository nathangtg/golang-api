package main

import (
	"fmt"
	"project/restful-api/cmd/api"
	"project/restful-api/db"

	"github.com/jackc/pgx"
)

func main() {

	// Run the Database
	db, err := db.NewPostgresDB(
		pgx.ConnConfig{
			Host:     "localhost",
			Port:     5432,
			Database: "restful_api",
			User:     "postgres",
			Password: "password",
		},
	)

	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	} else {
		fmt.Print("Database connected")
		db.Begin()
	}

	server := api.New(":8080", nil)
	if err := server.Start(); err != nil {
		fmt.Print(err.Error())
		panic(err)
	}
}