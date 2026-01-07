package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"todo-api/database"
	"todo-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("RAW register password: '%s' (len=%d)\n", user.Password, len(user.Password))

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Println("Error hashing password:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = database.Client.
		Database("todo_db").
		Collection("users").
		InsertOne(ctx, user)

	if err != nil {
		log.Println("Error inserting user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var creds struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
	if err:= json.NewDecoder(r.Body).Decode(&creds); err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx , cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	var user models.User
	err := database.Client.
		Database("todo_db").
		Collection("users").
		FindOne(ctx,bson.M{"email":creds.Email}).
		Decode(&user)
	if err != nil{
		log.Println("Error finding user:", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Println("db pass",user.Password)
	log.Println("login pass",creds.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(creds.Password))
	if err != nil {
		log.Println("Invalid password:", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	claims := jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_SECRET")
	signedToken , err := token.SignedString([]byte(secret))
	if err != nil{
		log.Println("Error signing token:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": signedToken,
	})

}

func MeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	userID, err := getUserID(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err = database.Client.
		Database("todo_db").
		Collection("users").
		FindOne(ctx, bson.M{"_id": userID}).
		Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
