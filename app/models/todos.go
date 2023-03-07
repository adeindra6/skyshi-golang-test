package models

import (
	"github.com/adeindra6/skyshi-golang-test/app/config"
	"gorm.io/gorm"
)

type Todos struct {
	gorm.Model
	ActivityGroupID uint   `gorm:"column:activity_group_id" json:"activity_group_id"`
	Title           string `gorm:"column:title" json:"title"`
	Priority        string `gorm:"column:priority" json:"priority"`
	IsActive        bool   `gorm:"column:is_active" json:"is_active"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Todos{})
}

func (t *Todos) CreateTodos() *Todos {
	var result Todos

	t.Priority = "very-high"
	db.Create(&t)
	db.Last(&result)

	return &result
}

func GetAllTodos() []Todos {
	var Todos []Todos
	db.Find(&Todos)
	return Todos
}

func GetTodoById(id int64) (*Todos, *gorm.DB) {
	var getTodo Todos
	db := db.Where("id = ?", id).Find(&getTodo)
	return &getTodo, db
}

func DeleteTodo(id int64) bool {
	var todos Todos
	var findTodos int64

	db.Model(&Todos{}).Where("id = ?", id).Count(&findTodos)
	if findTodos > 0 {
		db.Model(&Todos{}).Where("id = ?", id).Delete(&todos)
		return true
	} else {
		return false
	}
}
