package models

import (
	"fmt"
	"github.com/olusolaa/go-todo/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type  Todo struct {
	gorm.Model
	Title string `gorm:"" json:"title"`
}

func init()  {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Todo{})
	if err != nil {
		return 
	}
}

func (t *Todo) CreateTodo() *Todo {
	db.Create(&t)
	fmt.Println(t.ID)
	return t
}

func GetAllTodo(code string) []Todo {
	var Todos []Todo
	db.Where(fmt.Sprintf("title LIKE '%%%s%%'", code)).Find(&Todos)
	return Todos
}

func GetTodoById(ID int64) (*Todo, *gorm.DB)  {
	var getTodo Todo
	db := db.Where("ID=?", ID).Find(&getTodo)
	return &getTodo, db
}

func DeleteTodo(ID int64) Todo {
	var todo Todo
	db.Where("ID=?", ID).Delete(&todo)
	fmt.Println(&todo.ID)
	return todo
}

