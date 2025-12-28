package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	db_url := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), db_url)

	if err != nil {
		log.Fatal("Error while connecting to DB.")
	}

	DB = pool
	log.Println("Database connection successfull.")
}
