package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InnitializeDBService() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil && os.Getenv("ENV") != "heroku" {
		log.Fatalf("Error loading .env file")
	}

	dbcon, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil && os.Getenv("ENV") != "heroku" {
		panic(err)
	}

	// check the connection
	err = dbcon.Ping()

	if err != nil && os.Getenv("ENV") != "heroku" {
		panic(err)
	}

	fmt.Println("Successfully connected to DB!")
	return dbcon
}
