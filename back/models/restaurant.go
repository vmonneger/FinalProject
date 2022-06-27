package models

type Restaurant struct {
	Name string      `json:"name,omitempty" validate:"required"`
	Menu interface{} `json:"menu"`
}
