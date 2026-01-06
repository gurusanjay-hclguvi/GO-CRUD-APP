package main
import (
	"todo-api/handlers"
	"net/http"
)
func main(){
	http.HandleFunc("/todos/",handlers.Todoshandler)
	http.HandleFunc("/todos",handlers.Todoshandler)
	http.ListenAndServe(":8080",nil)
}