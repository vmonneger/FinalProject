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
)

func RestaurantHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var restaurant models.Restaurant
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
		if err := json.NewDecoder(r.Body).Decode(&restaurant); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.RestaurantResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&restaurant); validationErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.RestaurantResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		newRestaurant := models.Restaurant{
			Name: restaurant.Name,
			Menu: restaurant.Menu,
		}

		userIdConvert, _ := primitive.ObjectIDFromHex(userId.ID)
		fmt.Println(userIdConvert)
		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": userIdConvert}, bson.M{"$set": newRestaurant})
		fmt.Println(result)
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
			Data:    map[string]interface{}{"data": newRestaurant}}
		json.NewEncoder(w).Encode(response)
	}
}
