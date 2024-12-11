package routes

import (
	"github.com/gorilla/mux"
	"user_management_go/controllers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	//router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	//router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	//router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	//router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	return router
}
