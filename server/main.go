package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"todo-api/database"
	"todo-api/handlers"
	"todo-api/middleware"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI not set")
	}

	if err := database.ConnectToMongoDB(mongoURI); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")

	mux := http.NewServeMux()

	// Auth (public)
	mux.HandleFunc("/auth/register", handlers.RegisterHandler)
	mux.HandleFunc("/auth/login", handlers.LoginHandler)

	// Todos (protected)
	mux.HandleFunc("/todos", middleware.AuthMiddleware(handlers.Todoshandler))
	mux.HandleFunc("/todos/", middleware.AuthMiddleware(handlers.Todoshandler))
	mux.HandleFunc("/users/",middleware.AuthMiddleware(handlers.GetTodosByUserID))
	mux.HandleFunc("/auth/me/",middleware.AuthMiddleware(handlers.MeHandler))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(mux)))
}
