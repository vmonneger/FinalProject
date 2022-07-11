package controllers

// This controller manage all request about restaurant.

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

// Post essential restaurant info. Like Name, description...
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
			response := responses.RequestResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		// use the validator library to validate required fields
		if validationErr := validate.Struct(&restaurant); validationErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.RequestResponse{
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

		userIdConvert, _ := primitive.ObjectIDFromHex(userId.ID)

		_, err := userCollection.UpdateOne(ctx, bson.M{"_id": userIdConvert}, bson.M{"$set": newRestaurant})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.RequestResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusCreated)
		response := responses.RequestResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    newRestaurant}
		json.NewEncoder(w).Encode(response)
	}
}

// Get data of the restaurant.
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
			response := responses.RequestResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusCreated)
		response := responses.RequestResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    user}
		json.NewEncoder(w).Encode(response)
	}
}

// Post menu of the restaurant.
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
			response := responses.RequestResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		userIdConvert, _ := primitive.ObjectIDFromHex(userId.ID)

		_, err := userCollection.UpdateOne(ctx, bson.M{"_id": userIdConvert}, bson.D{primitive.E{Key: "$addToSet", Value: bson.D{primitive.E{Key: "menu", Value: bson.D{primitive.E{Key: "$each", Value: menu.Menu}}}}}})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.RequestResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusCreated)
		response := responses.RequestResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    menu.Menu}
		json.NewEncoder(w).Encode(response)
	}
}

// Post menu category of the restaurant.
func RestaurantCategoryPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var category models.Category
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
		if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.RequestResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		userIdConvert, _ := primitive.ObjectIDFromHex(userId.ID)

		_, err := userCollection.UpdateOne(ctx, bson.M{"_id": userIdConvert}, bson.M{"$set": bson.M{"category": category.Name}})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.RequestResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusCreated)
		response := responses.RequestResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    category.Name}
		json.NewEncoder(w).Encode(response)
	}
}

// Delete menu restaurant.
func RestaurantDeleteMenu() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var menuItem models.MenuItem

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

		if err := json.NewDecoder(r.Body).Decode(&menuItem); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.RequestResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		userId := token.Claims.(*services.Token)

		userIdConvert, _ := primitive.ObjectIDFromHex(userId.ID)

		fmt.Println(menuItem)

		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": userIdConvert}, bson.M{"$pull": bson.M{"menu": bson.M{"title": menuItem.Title, "description": menuItem.Description, "category": menuItem.Category}}})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.RequestResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusCreated)
		response := responses.RequestResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    result}
		json.NewEncoder(w).Encode(response)
	}
}

// Delete menu category.
func RestaurantDeleteCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var category models.Category
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

		if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.RequestResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		fmt.Println("tesssst")

		userId := token.Claims.(*services.Token)
		fmt.Println(category)

		userIdConvert, _ := primitive.ObjectIDFromHex(userId.ID)

		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": userIdConvert}, bson.M{"$pull": bson.M{"category": bson.M{"$in": category.Name}}})

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.RequestResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusCreated)
		response := responses.RequestResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    result}
		json.NewEncoder(w).Encode(response)
	}
}
