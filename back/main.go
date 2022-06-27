package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vmonneger/FinalProject/configs"
	"github.com/vmonneger/FinalProject/middleware"
	"github.com/vmonneger/FinalProject/routes"
)

func main() {
	router := mux.NewRouter()

	configs.ConnectDB()
	router.Use(middleware.IsAuthorized)
	// router.HandleFunc("/restaurant", test)
	// router.HandleFunc("/restaurant/{id}", test)

	routes.UserRoute(router)

	fmt.Printf("Starting server at port 6000\n")
	if err := http.ListenAndServe(":6000", router); err != nil {
		log.Fatal(err)
	}

}
