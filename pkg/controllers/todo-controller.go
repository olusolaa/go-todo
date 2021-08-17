package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/olusolaa/go-todo/pkg/models"
	"github.com/olusolaa/go-todo/pkg/util"
	"net/http"
	"strconv"
)

var NewTodo models.Todo

func CreateTodo(w http.ResponseWriter, r *http.Request)  {
	createTodo := &models.Todo{}
	util.ParseBody(r, createTodo)
	t:= createTodo.CreateTodo()
	res, _ :=json.Marshal(t)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	NewTodos := models.GetAllTodo(code)
	res, _ := json.Marshal(NewTodos)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}

func GetTodoById(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	todoId :=vars["todoId"]
	Id, err := strconv.ParseInt(todoId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	todo, _ := models.GetTodoById(Id)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func DeleteTodo(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	todoId :=vars["todoId"]
	Id, err := strconv.ParseInt(todoId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	todo := models.DeleteTodo(Id)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func UpdateTodo(w http.ResponseWriter, r *http.Request)  {
	updateTodo := &models.Todo{}
	util.ParseBody(r, updateTodo)
	vars := mux.Vars(r)
	todoId :=vars["todoId"]
	Id, err := strconv.ParseInt(todoId,0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	todo, db := models.GetTodoById(Id)
	if updateTodo.Title != ""{
		todo.Title = updateTodo.Title
	}
	db.Save(todo)
	res, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}