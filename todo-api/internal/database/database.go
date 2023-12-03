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

// Todo represents a ToDo item
type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

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

// GetTodos retrieves a list of all ToDo items from the database
func GetTodos() ([]Todo, error) {
	rows, err := db.Query("SELECT id, title, description FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
