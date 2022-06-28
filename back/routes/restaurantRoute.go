package routes

import (
	"github.com/gorilla/mux"
	"github.com/vmonneger/FinalProject/controllers"
)

func RestaurantRoute(router *mux.Router) {
	router.HandleFunc("/restaurant", controllers.RestaurantPost()).Methods("POST")
	router.HandleFunc("/restaurant", controllers.RestaurantGetOne()).Methods("GET")
}
