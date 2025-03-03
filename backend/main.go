package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type Response struct {
	Message string `json:"message"`
}

var db *sql.DB

func initDB() {
	var err error
	// Read the DATABASE_URL environment variable
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Connect to the database
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}

	log.Println("Successfully connected to the database!")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Example: Query the database
	var message string
	err := db.QueryRow("SELECT 'Hello from PostgreSQL!'").Scan(&message)
	if err != nil {
		http.Error(w, "Failed to query the database", http.StatusInternalServerError)
		return
	}

	response := Response{Message: message}
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Initialize the database connection
	initDB()
	defer db.Close()

	// Define routes
	http.HandleFunc("/hello", helloHandler)

	// Start the server
	log.Println("Backend is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
