package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vmonneger/FinalProject/configs"
	"github.com/vmonneger/FinalProject/middlewares"
	"github.com/vmonneger/FinalProject/routes"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST"},
	})

	router := mux.NewRouter()

	configs.ConnectDB()
	router.Use(mux.CORSMethodMiddleware(router), middlewares.IsAuthorized)

	routes.UserRoute(router)
	routes.RestaurantRoute(router)
	routes.PlaceRoute(router)

	fmt.Printf("Starting server at port 8000\n")

	if err := http.ListenAndServe(":8000", c.Handler((router))); err != nil {
		log.Fatal(err)
	}

}
