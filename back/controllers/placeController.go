package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/vmonneger/FinalProject/configs"
	"github.com/vmonneger/FinalProject/models"
	"github.com/vmonneger/FinalProject/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var placeCollection *mongo.Collection = configs.GetCollection(configs.DB, "places")

func PlacePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var place models.Place
		defer cancel()

		// Validate the request body
		if err := json.NewDecoder(r.Body).Decode(&place); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.RestaurantResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&place); validationErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.RestaurantResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		newPlace := models.Place{
			Id:      primitive.NewObjectID(),
			Name:    place.Name,
			User_id: place.User_id,
		}

		result, err := placeCollection.InsertOne(ctx, newPlace)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.RestaurantResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		var convertUsersId []primitive.ObjectID
		for _, element := range newPlace.User_id {
			id, _ := primitive.ObjectIDFromHex(element)
			convertUsersId = append(convertUsersId, id)
		}

		fmt.Printf("%+v\n", convertUsersId)

		resultUpdateUsers, err := userCollection.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": convertUsersId}}, bson.M{"$set": bson.M{"place_id": newPlace.Id}})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.RestaurantResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusCreated)
		response := responses.RestaurantResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"dataPlace": newPlace, "dataUserPlace": resultUpdateUsers, "mongodb": result}}
		json.NewEncoder(w).Encode(response)
	}
}

func PlaceGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		id, _ := primitive.ObjectIDFromHex(params["id"])
		project := bson.M{"password": 0, "email": 0}
		opts := options.Find().SetProjection(project)
		var users []models.User
		var place models.Place
		defer cancel()

		cursor, err := userCollection.Find(ctx, bson.M{"place_id": id}, opts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.RestaurantResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var result models.User
			err := cursor.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, result)
		}
		if err := cursor.Err(); err != nil {
			log.Fatal(err)
		}

		errPlace := placeCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&place)

		if errPlace != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusCreated)
		response := responses.RestaurantResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"dataRestaurant": users, "dataPlace": place}}
		json.NewEncoder(w).Encode(response)
	}
}
