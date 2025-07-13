package db

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/shubham-oop/helloWorld/config"
)

var DB *sqlx.DB

func Init() {
	var err error
	DB, err = sqlx.Connect("pgx", config.DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("Connected to PostgreSQL database")
}
