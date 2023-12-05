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

// GetTodoByID retrieves a single ToDo item by ID from the database
func GetTodoByID(id int) (*Todo, error) {
	var todo Todo
	err := db.QueryRow("SELECT id, title, description FROM todos WHERE id = $1", id).Scan(&todo.ID, &todo.Title, &todo.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Todo not found")
		}
		return nil, err
	}
	return &todo, nil
}

// InsertTodo inserts a new ToDo item into the database
func InsertTodo(todo Todo) error {
	_, err := db.Exec("INSERT INTO todos (title, description) VALUES ($1, $2)", todo.Title, todo.Description)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTodoByID updates a ToDo item in the database by ID
func UpdateTodoByID(id int, todo Todo) error {
	result, err := db.Exec("UPDATE todos SET title=$1, description=$2 WHERE id=$3", todo.Title, todo.Description, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No todo found with ID %d", id)
	}
	return nil
}
