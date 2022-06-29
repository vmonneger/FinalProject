package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/vmonneger/FinalProject/configs"
	"github.com/vmonneger/FinalProject/models"
	"github.com/vmonneger/FinalProject/responses"
	"github.com/vmonneger/FinalProject/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var placeCollection *mongo.Collection = configs.GetCollection(configs.DB, "places")

func PlacePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var place models.Place
		defer cancel()

		reqToken := r.Header.Get("Authorization")
		tokenString := strings.Split(reqToken, "Bearer ")[1]

		// Decode from the struct
		t := services.Token{}
		token, _ := jwt.ParseWithClaims(tokenString, &t, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return []byte(configs.EnvJwtSecret()), nil
		})

		userId := token.Claims.(*services.Token)

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
			Name: place.Name,
		}

		fmt.Println(newPlace)

		userIdConvert, _ := primitive.ObjectIDFromHex(userId.ID)

		result, err := placeCollection.UpdateOne(ctx, bson.M{"_id": userIdConvert}, bson.M{"$set": newPlace})

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
			Data:    map[string]interface{}{"data": newPlace, "mongodb": result}}
		json.NewEncoder(w).Encode(response)
	}
}
