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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RestaurantPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var restaurant models.Restaurant
		defer cancel()

		reqToken := r.Header.Get("Authorization")
		tokenString := strings.Split(reqToken, "Bearer ")[1]

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

		// use the validator library to validate required fields
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
			Name:        restaurant.Name,
			Description: restaurant.Description,
		}

		fmt.Println(newRestaurant)

		userIdConvert, _ := primitive.ObjectIDFromHex(userId.ID)

		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": userIdConvert}, bson.M{"$set": newRestaurant})

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
			Data:    map[string]interface{}{"data": newRestaurant, "mongodb": result}}
		json.NewEncoder(w).Encode(response)
	}
}

func RestaurantGetOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		reqToken := r.Header.Get("Authorization")
		tokenString := strings.Split(reqToken, "Bearer ")[1]

		t := services.Token{}
		token, _ := jwt.ParseWithClaims(tokenString, &t, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return []byte(configs.EnvJwtSecret()), nil
		})

		userId := token.Claims.(*services.Token)

		userIdConvert, _ := primitive.ObjectIDFromHex(userId.ID)
		project := bson.M{"password": 0, "email": 0}
		opts := options.FindOne().SetProjection(project)

		err := userCollection.FindOne(ctx, bson.M{"_id": userIdConvert}, opts).Decode(&user)

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
			Data:    map[string]interface{}{"data": user}}
		json.NewEncoder(w).Encode(response)
	}
}

func RestaurantMenuPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var menu models.Menu
		defer cancel()

		reqToken := r.Header.Get("Authorization")
		tokenString := strings.Split(reqToken, "Bearer ")[1]

		t := services.Token{}
		token, _ := jwt.ParseWithClaims(tokenString, &t, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return []byte(configs.EnvJwtSecret()), nil
		})

		userId := token.Claims.(*services.Token)

		// Validate the request body
		if err := json.NewDecoder(r.Body).Decode(&menu); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		newMenu := models.Menu{
			Menu: menu.Menu,
		}

		userIdConvert, _ := primitive.ObjectIDFromHex(userId.ID)
		fmt.Println(newMenu)

		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": userIdConvert}, bson.M{"$set": newMenu})

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
			Data:    map[string]interface{}{"data": newMenu, "mongodb": result}}
		json.NewEncoder(w).Encode(response)
	}
}
