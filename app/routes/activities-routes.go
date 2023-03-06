package routes

import (
	"github.com/adeindra6/skyshi-golang-test/app/controllers"
	"github.com/gorilla/mux"
)

var RegisterActivitiesRoutes = func(router *mux.Router) {
	router.HandleFunc("/activity-groups", controllers.CreateActivities).Methods("POST")
	router.HandleFunc("/activity-groups", controllers.GetActivities).Methods("GET")
	router.HandleFunc("/activity-groups/{activity_id}", controllers.GetActivityById).Methods("GET")
	router.HandleFunc("/activity-groups/{activity_id}", controllers.UpdateActivities).Methods("PATCH")
	router.HandleFunc("/activity-groups/{activity_id}", controllers.DeleteActivity).Methods("DELETE")
}
