package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	// We put the _ in the front so that the import doesnot disapper when saving
	_ "github.com/lib/pq"
)

func main() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read the connection details from environment variables
	host := os.Getenv("PG_HOST")

	// Convert port to an integer
	portStr := os.Getenv("PG_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Error converting port to integer: %v", err)
	}

	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DBNAME")

	// Construct the connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connect to the PostgreSQL database (always check errors)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
}
