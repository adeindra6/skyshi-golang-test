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

var activites models.Activities

type SuccessMessageActivities struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    models.Activities `json:"data"`
}

type ArrSuccessMessageActivities struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Data    []models.Activities `json:"data"`
}

type DeleteMessageActivities struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrMessageActivities struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Code    int64  `json:"code"`
}

func CreateActivities(w http.ResponseWriter, r *http.Request) {
	CreateActivities := &models.Activities{}
	utils.ParseBody(r, CreateActivities)
	a := CreateActivities.CreateActivities()

	_, err := json.Marshal(a)
	if err != nil {
		err_msg := ErrMessageActivities{
			Status:  "ERROR",
			Message: "Error while creating new activities",
			Code:    http.StatusInternalServerError,
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	success_msg := SuccessMessageActivities{
		Status:  "Success",
		Message: "Success",
		Data:    *a,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func GetActivities(w http.ResponseWriter, r *http.Request) {
	activities := models.GetAllActitivities()

	_, err := json.Marshal(activities)
	if err != nil {
		err_msg := ErrMessageActivities{
			Status:  "ERROR",
			Message: "Error when fetching all activities",
			Code:    http.StatusInternalServerError,
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	success_msg := ArrSuccessMessageActivities{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func GetActivityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	activity_id := vars["activity_id"]

	id, err := strconv.ParseInt(activity_id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	activity, _ := models.GetActivityById(id)
	_, err = json.Marshal(activity)
	if err != nil {
		err_msg := ErrMessageActivities{
			Status:  "ERROR",
			Message: "Error when fetching activity",
			Code:    http.StatusInternalServerError,
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	success_msg := SuccessMessageActivities{
		Status:  "Success",
		Message: "Success",
		Data:    *activity,
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func UpdateActivities(w http.ResponseWriter, r *http.Request) {
	var updateActivity = &models.Activities{}
	utils.ParseBody(r, updateActivity)
	vars := mux.Vars(r)
	activity_id := vars["activity_id"]

	id, err := strconv.ParseInt(activity_id, 0, 0)
	if err != nil {
		err_msg := ErrMessageActivities{
			Status:  "ERROR",
			Message: "Error when updating activity",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	activity, db := models.GetActivityById(id)
	if updateActivity.Title != "" {
		activity.Title = updateActivity.Title
	}

	db.Save(&activity)
	_, err = json.Marshal(activity)
	if err != nil {
		fmt.Println("Error while parsing!!!")
	}

	success_msg := SuccessMessageActivities{
		Status:  "Success",
		Message: "Success",
		Data:    *activity,
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}

func DeleteActivity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	activity_id := vars["activity_id"]

	id, err := strconv.ParseInt(activity_id, 0, 0)
	if err != nil {
		err_msg := ErrMessageActivities{
			Status:  "ERROR",
			Message: "Error when deleting activity",
			Code:    http.StatusInternalServerError,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err_msg)
	}

	successDeleted := models.DeleteActivity(id)
	var success_msg DeleteMessageActivities
	if successDeleted {
		success_msg = DeleteMessageActivities{
			Status:  "Success",
			Message: fmt.Sprintf("Success Deleted id: %d", id),
		}
	} else {
		success_msg = DeleteMessageActivities{
			Status:  "Not Found",
			Message: fmt.Sprintf("Activity with ID %d Not Found", id),
		}
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success_msg)
}
