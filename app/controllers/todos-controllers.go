package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/adeindra6/skyshi-golang-test/app/models"
	"github.com/adeindra6/skyshi-golang-test/app/utils"
	"github.com/gorilla/mux"
)

var todos models.Todos

type SuccessMessageTodos struct {
	Status  string       `json:"status"`
	Message string       `json:"message"`
	Data    models.Todos `json:"data"`
}

type ArrSuccessMessageTodos struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Data    []models.Todos `json:"data"`
}

type DeleteMessageTodos struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrMessageTodos struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

func CreateTodos(w http.ResponseWriter, r *http.Request) {
	CreateTodos := &models.Todos{}
	utils.ParseBody(r, CreateTodos)
	t := CreateTodos.CreateTodos()

	_, err := json.Marshal(t)
	if err != nil {
		err_msg := ErrMessageActivities{
			Status:  "ERROR",
			Message: "Error while creating new todos",
			Code:    http.StatusInternalServerError,
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	success_msg := SuccessMessageTodos{
		Status:  "Success",
		Message: "Success",
		Data:    *t,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := models.GetAllTodos()

	_, err := json.Marshal(todos)
	if err != nil {
		err_msg := ErrMessageTodos{
			Status:  "ERROR",
			Message: "Error when fetching all todos",
			Code:    http.StatusInternalServerError,
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	success_msg := ArrSuccessMessageTodos{
		Status:  "Success",
		Message: "Success",
		Data:    todos,
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todo_id := vars["todo_id"]

	id, err := strconv.ParseInt(todo_id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	todo, _ := models.GetTodoById(id)
	_, err = json.Marshal(todo)
	if err != nil {
		err_msg := ErrMessageTodos{
			Status:  "ERROR",
			Message: "Error when fetching todo",
			Code:    http.StatusInternalServerError,
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	success_msg := SuccessMessageTodos{
		Status:  "Success",
		Message: "Success",
		Data:    *todo,
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func UpdateTodos(w http.ResponseWriter, r *http.Request) {
	var updateTodo = &models.Todos{}
	utils.ParseBody(r, updateTodo)
	vars := mux.Vars(r)
	todo_id := vars["todo_id"]

	id, err := strconv.ParseInt(todo_id, 0, 0)
	if err != nil {
		err_msg := ErrMessageTodos{
			Status:  "ERROR",
			Message: "Error when updating todo",
			Code:    http.StatusInternalServerError,
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	todo, db := models.GetTodoById(id)
	if updateTodo.Title != "" {
		todo.Title = updateTodo.Title
	}
	if updateTodo.Priority != "" {
		todo.Priority = updateTodo.Priority
	}
	if updateTodo.IsActive {
		todo.IsActive = true
	} else {
		todo.IsActive = false
	}

	db.Save(&todo)
	_, err = json.Marshal(todo)
	if err != nil {
		fmt.Println("Error while parsing!!!")
	}

	success_msg := SuccessMessageTodos{
		Status:  "Success",
		Message: "Success",
		Data:    *todo,
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todo_id := vars["todo_id"]

	id, err := strconv.ParseInt(todo_id, 0, 0)
	if err != nil {
		err_msg := ErrMessageTodos{
			Status:  "ERROR",
			Message: "Error when deleting todo",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	successDeleted := models.DeleteTodo(id)
	var success_msg DeleteMessageTodos
	if successDeleted {
		success_msg = DeleteMessageTodos{
			Status:  "Success",
			Message: fmt.Sprintf("Success Deleted id: %d", id),
		}
	} else {
		success_msg = DeleteMessageTodos{
			Status:  "Not Found",
			Message: fmt.Sprintf("Todo with ID %d Not Found", id),
		}
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}
