package routes

import (
	"github.com/gorilla/mux"
	"github.com/vmonneger/FinalProject/controllers"
)

func UserRoute(router *mux.Router) {
	router.HandleFunc("/auth/signin", controllers.UserSignIn()).Methods("POST")
	router.HandleFunc("/auth/login", controllers.UserLogin()).Methods("POST")
	router.HandleFunc("/users", controllers.UserGetAll()).Methods("GET")
}
