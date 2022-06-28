package models

type Restaurant struct {
	Name        string `json:"name,omitempty" validate:"required" bson:"name,omitempty"`
	Description string `json:"description,omitempty" validate:"required" bson:"description,omitempty"`
}
