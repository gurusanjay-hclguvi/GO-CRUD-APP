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
		origin := os.Getenv("CLIENT_ORIGIN")
		if origin == "" {
			origin = "*"
		}

		w.Header().Set("Access-Control-Allow-Origin", origin)
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
	// Load .env only for local dev
	if os.Getenv("ENV") != "production" {
		_ = godotenv.Load()
	}

	// ---- Database ----
	mongoURI := os.Getenv("DB_URL")
	if mongoURI == "" {
		log.Fatal("DB_URL not set")
	}

	if err := database.ConnectToMongoDB(mongoURI); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Println("Connected to MongoDB")

	// ---- Router ----
	mux := http.NewServeMux()

	// Auth (public)
	mux.HandleFunc("/auth/register", handlers.RegisterHandler)
	mux.HandleFunc("/auth/login", handlers.LoginHandler)

	// Protected routes
	mux.HandleFunc("/todos", middleware.AuthMiddleware(handlers.Todoshandler))
	mux.HandleFunc("/todos/", middleware.AuthMiddleware(handlers.Todoshandler))
	mux.HandleFunc("/users/", middleware.AuthMiddleware(handlers.GetTodosByUserID))
	mux.HandleFunc("/auth/me/", middleware.AuthMiddleware(handlers.MeHandler))

	// ---- Server ----
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, corsMiddleware(mux)))
}
