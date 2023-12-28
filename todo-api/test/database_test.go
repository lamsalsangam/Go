// tests/database_test.go
package tests

import (
	"testing"

	"todo-api/internal/database"
)

func TestGetTodos(t *testing.T) {
	// Initialize the test database
	err := database.InitDB()
	if err != nil {
		t.Fatal(err)
	}

	// Test fetching todos
	todos, err := database.GetTodos()
	if err != nil {
		t.Errorf("Error fetching todos: %v", err)
	}

	// Check if todos is not nil
	if todos == nil {
		t.Error("Expected todos, got nil")
	}
}

// Additional tests for other database functions can be added here
