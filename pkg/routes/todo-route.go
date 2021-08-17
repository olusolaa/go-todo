package routes

import (
	"github.com/gorilla/mux"
	"github.com/olusolaa/go-todo/pkg/controllers"
)

var TodoRoutes = func(router *mux.Router){
	router.HandleFunc("/todo", controllers.CreateTodo).Methods("POST")
	router.HandleFunc("/todo", controllers.GetTodo).Methods("GET")
	router.HandleFunc("/todo/{todoId}", controllers.GetTodoById).Methods("GET")
	router.HandleFunc("/todo/{todoId}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{todoId}", controllers.DeleteTodo).Methods("DELETE")
}