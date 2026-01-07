package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"todo-api/database"
	"todo-api/handlers"
	
)

// CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Allow Nuxt dev server
		w.Header().Set("Access-Control-Allow-Origin", "localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// loading .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set")
	}
	db_err := database.ConnectToMongoDB(mongoURI)
	if db_err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Println("Connected To Mongo DB")
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", handlers.Todoshandler)
	mux.HandleFunc("/todos/", handlers.Todoshandler)
	mux.HandleFunc("/auth/register",handlers.RegisterHandler)
	mux.HandleFunc("/auth/login",handlers.LoginHandler)

	http.ListenAndServe(":8080", corsMiddleware(mux))
}
