package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yourusername/blog/internal/interfaces"
	"github.com/yourusername/blog/pkg/web"
)

func main() {
	router := mux.NewRouter()

	// Initialize database
	db := interfaces.NewPostgreSQLDB("your_postgres_connection_string")
	defer db.Close()

	// Initialize use cases
	postInteractor := interfaces.NewBlogPostInteractor(db)

	// Initialize handlers
	postHandler := web.NewBlogPostHandler(postInteractor)

	// Define routes
	router.HandleFunc("/posts", postHandler.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", postHandler.GetPost).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", router))
}
