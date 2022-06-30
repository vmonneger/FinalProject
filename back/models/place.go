package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Place struct {
	Id      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" validate:"required" bson:"_name"`
	User_id []string           `json:"user_id" validate:"required" bson:"user_id"`
}
