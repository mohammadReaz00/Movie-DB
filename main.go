package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"mohammadreaz.com/moviedb/handlers"
	"mohammadreaz.com/moviedb/logger"
)

func initializeLogger() *logger.Logger {
	logInstance, err := logger.NewLogger("move.log")
	if err != nil {
		log.Fatalf("Failed to initialize logger %v", err)
	}
	defer logInstance.Close()
	return logInstance
}

func main() {

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or failed to load: %v", err)
	}

	// Initialize logger
	logInstance := initializeLogger()

	// Database connection
	dbConnStr := os.Getenv("DATABASE_URL")
	if dbConnStr == "" {
		log.Fatalf("DATABASE_URL not set in environment")
	}
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	movieHandler := handlers.MovieHandler{}
	// Set up routes
	http.HandleFunc("/api/movies/top/", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/random/", movieHandler.GetRandomMovies)

	//handler for static files
	http.Handle("/", http.FileServer(http.Dir("public")))

	const addr = ":8080"
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Server Failed: %v", err)
		logInstance.Error("Server Failed", err)
	}

}
