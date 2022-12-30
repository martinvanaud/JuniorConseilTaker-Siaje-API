package database

import (
	"database/sql"

	"fmt"
	"log"
	"os"
)

func Configure() (*sql.DB, error) {
	client, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_NAME"),
			os.Getenv("POSTGRES_PASSWORD"),
		))

	if err != nil {
		log.Fatalf("postgresql: could not proerly setup instance %v", err)
	} else {
		fmt.Println("Postgres database is ready")
	}

	return client, err
}
