package models

type Category struct {
	Name []string `json:"name" validate:"required" bson:"name"`
}
