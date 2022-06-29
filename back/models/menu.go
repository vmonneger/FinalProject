package models

type Menu struct {
	Menu interface{} `json:"menu,omitempty" validate:"required" bson:"menu,omitempty"`
}
