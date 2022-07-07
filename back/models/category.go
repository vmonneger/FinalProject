package models

// type Category struct {
// 	Id      primitive.ObjectID `json:"id" bson:"_id"`
// 	Name    string             `json:"name" validate:"required" bson:"_name"`
// 	User_id []string           `json:"user_id" validate:"required" bson:"user_id"`
// }

type Category struct {
	Name []string `json:"name" validate:"required" bson:"name"`
}
