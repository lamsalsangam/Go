package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	mu sync.Mutex
)

const (
	host     = "localhost"
	port     = 5432
	user     = "your_username"
	password = "your_password"
	dbname   = "your_dbname"
)

// InitDB initializes the database connection
func InitDB() error {
	mu.Lock()
	defer mu.Unlock()

	if db != nil {
		return nil
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("Connected to the database")
	return nil
}

// GetDB returns the initialized database connection
func GetDB() (*sql.DB, error) {
	if db == nil {
		return nil, fmt.Errorf("databse not initialized")
	}
	return db, nil
}
