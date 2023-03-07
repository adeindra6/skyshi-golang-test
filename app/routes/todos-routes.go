package routes

import (
	"github.com/adeindra6/skyshi-golang-test/app/controllers"
	"github.com/gorilla/mux"
)

var RegisterTodosRoutes = func(router *mux.Router) {
	router.HandleFunc("/todo-items", controllers.CreateTodos).Methods("POST")
	router.HandleFunc("/todo-items", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/todo-items/{todo_id}", controllers.GetTodoById).Methods("GET")
	router.HandleFunc("/todo-items/{todo_id}", controllers.UpdateTodos).Methods("PATCH")
	router.HandleFunc("/todo-items/{todo_id}", controllers.DeleteTodo).Methods("DELETE")
}
