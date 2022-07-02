package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserAccount struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Email    string             `json:"email,omitempty" validate:"required" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" validate:"required" bson:"password,omitempty"`
}
