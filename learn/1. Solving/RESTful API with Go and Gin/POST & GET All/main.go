package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// +++++++++++++++++++++++++++++++++++++++++++++++

// Initialize a Gin router using Default.

// Use the GET function to associate the GET HTTP method and /albums path with a handler function.

// Note that you're passing the name of the getAlbums function. This is different from passing the result of the function, which you would do by passing getAlbums() (note the parenthesis).

// Use the Run function to attach the router to an http.Server and start the server.

// Associate the POST method at the /albums path with the postAlbums function.

// With Gin, you can associate a handler with an HTTP method-and-path combination. In this way, you can separately route requests sent to a single path based on the method the client is using.

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// +++++++++++++++++++++++++++++++++++++++++++++++

// Write a getAlbums function that takes a gin.Context parameter. Note that you could have given this function any name – neither Gin nor Go require a particular function name format.

// gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. (Despite the similar name, this is different from Go's built-in context package.)

// Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.

// The function's first argument is the HTTP status code you want to send to the client. Here, you're passing the StatusOK constant from the net/http package to indicate 200 OK.

// Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON. In practice, the indented form is much easier to work with when debugging and the size difference is usually small.

// getAlbums responds with the list of all albums as JSON.

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// +++++++++++++++++++++++++++++++++++++++++++++++

// Use Context.BindJSON to bind the request body to newAlbum.

// Append the album struct initialized from the JSON to the albums slice.

// Add a 201 status code to the response, along with JSON representing the album you added.

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// +++++++++++++++++++++++++++++++++++++++++++++++
// TERMINAL

// Run the command curl http://localhost:8080/albums to get the result

// curl http://localhost:8080/albums \
//     --include --header \
//     "Content-Type: application/json" \
//     --request "POST" --data \
//     '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

// curl http://localhost:8080/albums \
//     --header \
//     "Content-Type: application/json" \
//     --request "GET"

// curl -X GET http://localhost:8080/albums -H "Content-Type: application/json"
// curl --request GET http://localhost:8080/albums -header "Content-Type: application/json"
