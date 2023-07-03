package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/gorilla/mux"

	"my-book-app/api"
	"my-book-app/storage"
)

func main() {
	// Connect to the PostgreSQL database
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Create the book repository
	bookRepo := storage.NewPostgresBookRepository(db)

	// Initialize the router
	router := mux.NewRouter()

	// Register book routes
	bookHandler := api.NewBookHandler(*bookRepo)
	router.HandleFunc("/books/create", bookHandler.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", bookHandler.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", bookHandler.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", bookHandler.DeleteBook).Methods("DELETE")

	// Start the server
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func connectToDatabase() (*sql.DB, error) {
	godotenv.Load()
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	ip := os.Getenv("IP")
	database := os.Getenv("DATABASE")
	// Set up the PostgreSQL connection parameters
	connStr := "postgres://"+username+":"+password+"@"+ip+":5432/"+database+"?sslmode=disable"
	
	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")

	err = createBooksTable(db)
	if err != nil {
		log.Fatal("Failed to create books table:", err)
	}

	return db, nil
}

func createBooksTable(db *sql.DB) error {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS books (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			author VARCHAR(255) NOT NULL,
			genre VARCHAR(255) NOT NULL,
			year INT NOT NULL
		)
	`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}
