package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"todo-api/models"
)
var todos []models.Todo
var nextID = 1
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
func CreateTodo(w http.ResponseWriter , r *http.Request) {
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo.ID = nextID
	nextID++
	todos = append(todos, todo)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
	
}
func GetTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(todos)

}
func GetTodoByID (w http.ResponseWriter , r *http.Request){
	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path,"/todos/")
	id , err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for _, todo := range todos {
		if todo.ID == id {
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
func UpdateTodo(w http.ResponseWriter , r *http.Request){
	if r.Method != http.MethodPut{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path,"/todos/")
	id , err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var updatedTodo models.Todo
	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i , todo := range todos{
		if todo.ID == id {
			updatedTodo.ID = id
			todos[i] = updatedTodo
			w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedTodo)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
func DeleteTodo(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	idStr := strings.TrimPrefix(r.URL.Path,"/todos/")
	id , err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	for i , todo := range todos{
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}