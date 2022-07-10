package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Email       string             `json:"email,omitempty" validate:"required" bson:"email,omitempty"`
	Password    string             `json:"password,omitempty" validate:"required" bson:"password,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required" bson:"name,omitempty"`
	Description string             `json:"description,omitempty" validate:"required" bson:"description,omitempty"`
	Menu        []interface{}      `json:"menu,omitempty" validate:"required" bson:",inline,omitempty"`
	Category    []string           `json:"category,omitempty" bson:"category,omitempty"`
	Place_id    string             `json:"place_id" bson:"place_id"`
}
