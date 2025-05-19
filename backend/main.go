package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AB-Rhman/simple-go/db"
	"github.com/AB-Rhman/simple-go/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	db.InitDB()

	// Create router
	r := mux.NewRouter()

	// Create handler with database
	h := handlers.NewHandler(db.DB)

	// Define routes
	r.HandleFunc("/api/tasks", h.GetTasks).Methods("GET")
	r.HandleFunc("/api/tasks", h.CreateTask).Methods("POST")
	r.HandleFunc("/api/tasks/{id}", h.DeleteTask).Methods("DELETE")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
