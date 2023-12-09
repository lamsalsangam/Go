package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"todo-api/internal/database"
)

// GetTodos retrieves a list of all ToDo items
func GetTodos(w http.ResponseWriter, r *http.Request) {
	err := database.InitDB()
	if err != nil {
		log.Printf("Error initializing the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	todos, err := database.GetTodos()
	if err != nil {
		log.Printf("Error fetching todos: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, todos)
}

// GetTodo retrieves a single ToDo item by ID
func GetTodo(w http.ResponseWriter, r *http.Request) {
	err := database.InitDB()
	if err != nil {
		log.Printf("Error initializing the database:%v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	idStr := extractIDFromURL(r)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	todo, err := database.GetTodoByID(id)
	if err != nil {
		log.Printf("Error fetching todo: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, todo)
}

// respondWithJSON sends a JSON response with the specified status code and data
func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// extractIDFromURL extracts the todo ID from the URL path
func extractIDFromURL(r *http.Request) string {
	return r.URL.Path[len("/todos/"):]
}
