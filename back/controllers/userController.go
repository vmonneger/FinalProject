package controllers

// This controller manage all about user. Like register, login, admin...

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

// User signin.
func UserSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.UserAccount
		defer cancel()

		// Validate the request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.RequestResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.RequestResponse{
				Status:  http.StatusBadRequest,
				Message: "error",
				Data:    map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		opt := options.Index().SetUnique(true)
		index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

		if _, err := userCollection.Indexes().CreateOne(ctx, index); err != nil {
			services.ServerErrResponse(err.Error(), w)
		}

		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		newUser := models.UserAccount{
			Id:       primitive.NewObjectID(),
			Email:    user.Email,
			Password: string(hashPassword),
		}

		result, err := userCollection.InsertOne(ctx, newUser)

		if err != nil {
			services.ServerErrResponse(err.Error(), w)
			return
		}

		w.WriteHeader(http.StatusCreated)
		response := responses.RequestResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"data": result}}
		json.NewEncoder(w).Encode(response)
	}
}

// User login.
func UserLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User
		var dbUser models.User

		json.NewDecoder(r.Body).Decode(&user)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)
		defer cancel()

		if err != nil {
			services.ServerErrResponse(err.Error(), w)
			return
		}

		userPass := []byte(user.Password)
		dbPass := []byte(dbUser.Password)

		passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

		if passErr != nil {
			services.UnauthorizedErrResponse("Wrong Password", w)
			return
		}

		jwtToken, err := services.CreateToken(dbUser.Id.Hex(), user.Email)

		if err != nil {
			services.ServerErrResponse(err.Error(), w)
			return
		}

		response := map[string]interface{}{"token": jwtToken}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// Get all users.
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
			services.ServerErrResponse(err.Error(), w)
			return
		}

		w.WriteHeader(http.StatusCreated)
		response := responses.RequestResponse{
			Status:  http.StatusCreated,
			Message: "success",
			Data:    map[string]interface{}{"data": users}}
		json.NewEncoder(w).Encode(response)
	}
}

// Create a reference in the DB to link restaurant to a place.
func UserAddPlace(userId [1]string, placeId string) *mongo.UpdateResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := userCollection.UpdateMany(ctx, bson.M{"_id": bson.M{"$in": userId}}, bson.M{"$set": bson.M{"place_id": placeId}})

	if err != nil {
		log.Fatal(err)
	}
	return result
}
