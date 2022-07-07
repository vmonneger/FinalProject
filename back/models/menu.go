package models

import "go.mongodb.org/mongo-driver/bson"

type Menu struct {
	Menu []bson.M `json:"menu,omitempty" validate:"required" bson:"menu,omitempty"`
}
