package models

type Category struct {
	Category []string `json:"category,omitempty" validate:"required" bson:"category,omitempty"`
}
