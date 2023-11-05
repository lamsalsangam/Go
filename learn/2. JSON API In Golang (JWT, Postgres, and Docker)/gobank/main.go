package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// func main() {
// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
// 	// })
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Welcome to my website!")
// 	})

// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println("Your server is running")
// 	fmt.Println()
// 	fmt.Println("localhost:\t\t\tlocalhost:8080")
// 	fmt.Println()
// 	fmt.Println()
// 	fmt.Println("Press Ctrl + C to quit:")
// 	fmt.Println()

// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		// The package comes with the function mux.Vars(r) which takes the http.Request as parameter and returns a map of the segments.
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)
}
