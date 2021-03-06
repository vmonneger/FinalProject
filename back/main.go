package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vmonneger/FinalProject/configs"
	"github.com/vmonneger/FinalProject/middlewares"
	"github.com/vmonneger/FinalProject/routes"
	"github.com/vmonneger/FinalProject/services"
)

func main() {
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	})

	router := mux.NewRouter()

	configs.ConnectDB()
	router.Use(mux.CORSMethodMiddleware(router), middlewares.IsAuthorized, services.SetHeader)

	routes.UserRoute(router)
	routes.RestaurantRoute(router)
	routes.PlaceRoute(router)

	if err := http.ListenAndServe(configs.EnvServerPort(), c.Handler(router)); err != nil {
		log.Fatal(err)
	}

}
