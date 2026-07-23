package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"event-rest-api/internal/database"
	"event-rest-api/internal/env"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
}

func main() {
	db, err := sql.Open("sqlite3", "./datadb")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	models := database.NewModels(db)
	app := &application{
		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "SOME-SECRET-123456"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}