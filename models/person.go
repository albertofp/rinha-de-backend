package models

type Person struct {
	Nickname  string   `json:"nickname" bson:"nickname" validate:"required,max=32"`
	Name      string   `json:"name" bson:"nickname" validate:"required,max=100"`
	Birthdate string   `json:"birthdate" bson:"birthdate" validate:"required,2023-08-15"`
	Stack     []string `json:"stack" bson:"stack" validate:"dive,max=32"`
}
