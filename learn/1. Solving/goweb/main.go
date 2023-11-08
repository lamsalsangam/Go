package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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

	// // check db
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Connected!")

	// {} represents the block of code which will have its own scope
	// Creating the TABLE initally
	{
		query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at TIMESTAMPTZ
		);`

		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	{ //Insert a new user
		username := "Happy"
		password := "Home"
		createdAt := time.Now()

		_, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}
	}

	{ // Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = $1"
		err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{ // Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}
}
