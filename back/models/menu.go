package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Menu struct {
	Menu []interface{} `json:"menu,omitempty" validate:"required" bson:"menu,omitempty"`
}

type MenuItem struct {
	Id          primitive.ObjectID `json:"id" bson:"_id"`
	Description string             `json:"description" bson:"description"`
	Title       string             `json:"title" bson:"title"`
	Category    string             `json:"category" bson:"category"`
	Price       int                `json:"price" bson:"price"`
}
