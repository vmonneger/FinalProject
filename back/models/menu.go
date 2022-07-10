package models

type Menu struct {
	Menu []interface{} `json:"menu,omitempty" validate:"required" bson:"menu,omitempty"`
}

type MenuItem struct {
	Description string `json:"description" bson:"description"`
	Title       string `json:"title" bson:"title"`
	Category    string `json:"category" bson:"category"`
}
