// cmd/todo-api/main.go
package main

import (
	"log"
	"net/http"

	"todo-api/internal/database"
	"todo-api/internal/handlers"
	"todo-api/pkg/middleware"
)

func main() {
	// Initialize the database
	err := database.InitDB()
	if err != nil {
		log.Fatal("Error initializing the database:", err)
	}

	// Set up HTTP server and routes
	mux := http.NewServeMux()

	// Register middleware
	mux.Handle("/", middleware.LoggingMiddleware(http.HandlerFunc(handlers.GetTodos)))

	// Register routes
	mux.HandleFunc("/todos", handlers.CreateTodo)
	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTodo(w, r)
		case http.MethodPut:
			handlers.UpdateTodo(w, r)
		case http.MethodDelete:
			handlers.DeleteTodo(w, r)
		}
	})

	// Start the server
	port := "8080"
	log.Println("Server listening on port", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
