package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Place struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"_name"`
}
