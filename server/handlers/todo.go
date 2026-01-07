package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"context"
	"time"
	"todo-api/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"todo-api/models"
)
var todos []models.Todo
func Todoshandler(w http.ResponseWriter , r *http.Request){
	if r.URL.Path == "/todos" || r.URL.Path == "/todos/"{ 
	switch r.Method {
	case http.MethodPost:
		CreateTodo(w,r)
	case http.MethodGet:
		GetTodos(w,r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	return
	}
	if strings.HasPrefix(r.URL.Path,"/todos/"){
		switch r.Method {
		case http.MethodGet:
			GetTodoByID (w,r)
		case http.MethodPut:
			UpdateTodo (w,r)
		case http.MethodDelete:
			DeleteTodo (w,r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
// Creating a new todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := getUserID(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	todo.ID = primitive.NewObjectID()
	todo.UserID = userID

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = database.Client.
		Database("todo_db").
		Collection("todos").
		InsertOne(ctx, todo)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
// Getting all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := database.Client.
		Database("todo_db").
		Collection("todos").
		Find(ctx, bson.M{"userId": userID})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var todos []models.Todo
	if err := cursor.All(ctx, &todos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
// Getting a todo by ID of a user
func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	// extract todo id from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	todoID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get userId from JWT context
	userID, err := getUserID(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var todo models.Todo
	err = database.Client.
		Database("todo_db").
		Collection("todos").
		FindOne(ctx, bson.M{
			"_id":    todoID,
			"userId": userID,
		}).
		Decode(&todo)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
// Get all todos of a user
func GetTodosByUserID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	idStr = strings.TrimSuffix(idStr, "/todos")

	userID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := database.Client.
		Database("todo_db").
		Collection("todos").
		Find(ctx, bson.M{"userId": userID})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var todos []models.Todo
	if err := cursor.All(ctx, &todos); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// Updating a TODO
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	todoID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := getUserID(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var update models.Todo
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := database.Client.
		Database("todo_db").
		Collection("todos").
		UpdateOne(
			ctx,
			bson.M{"_id": todoID, "userId": userID},
			bson.M{"$set": bson.M{
				"title":     update.Title,
				"completed": update.Completed,
			}},
		)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
// Deleting a TODO
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	todoID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := getUserID(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := database.Client.
		Database("todo_db").
		Collection("todos").
		DeleteOne(ctx, bson.M{
			"_id":    todoID,
			"userId": userID,
		})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
