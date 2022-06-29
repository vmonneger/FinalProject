package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/vmonneger/FinalProject/configs"
	"github.com/vmonneger/FinalProject/models"
	"github.com/vmonneger/FinalProject/responses"
	"github.com/vmonneger/FinalProject/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

// Validator to check require fields
var validate = validator.New()

func UserSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		// Validate the request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		newUser := models.User{
			Id:       primitive.NewObjectID(),
			Email:    user.Email,
			Password: string(hashPassword),
		}

		result, err := userCollection.InsertOne(ctx, newUser)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusCreated)
		response := responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"data": result}}
		json.NewEncoder(w).Encode(response)
	}
}

func UserLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User
		var dbUser models.User

		json.NewDecoder(r.Body).Decode(&user)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)
		defer cancel()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		userPass := []byte(user.Password)
		dbPass := []byte(dbUser.Password)

		passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

		if passErr != nil {
			log.Println(passErr)
			response := responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		jwtToken, err := services.CreateToken(dbUser.Id.Hex(), user.Email)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusCreated)
		response := responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"data": jwtToken}}
		json.NewEncoder(w).Encode(response)
	}
}

func UserGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var users []models.User

		project := bson.M{"name": 1, "email": 1}
		opts := options.Find().SetProjection(project)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		cursor, err := userCollection.Find(ctx, bson.M{}, opts)
		defer cancel()
		if err != nil {
			log.Fatal(err)
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

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		w.WriteHeader(http.StatusCreated)
		response := responses.UserResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"data": users}}
		json.NewEncoder(w).Encode(response)
	}
}
