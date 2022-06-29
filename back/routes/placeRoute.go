package routes

import (
	"github.com/gorilla/mux"
	"github.com/vmonneger/FinalProject/controllers"
)

func PlaceRoute(router *mux.Router) {
	router.HandleFunc("/admin/place", controllers.PlacePost()).Methods("POST")
}
