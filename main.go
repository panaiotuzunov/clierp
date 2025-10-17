package main

import (
	"clierp/internal/database"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type State struct {
	db *database.Queries
}

func main() {
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Opening DB connection failed")
	}
	stateStruct := State{
		db: database.New(db),
	}
	startRepl(&stateStruct)
}
