package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func InnitializeDBService() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		// log.Fatalf("Error loading .env file")
	}

	dbcon, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		// panic(err)
	}

	// check the connection
	err = dbcon.Ping()

	if err != nil {
		// panic(err)
	}

	fmt.Println("Successfully connected to DB!")
	return dbcon
}
