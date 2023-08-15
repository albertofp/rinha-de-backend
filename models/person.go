package models

import "github.com/google/uuid"

type Person struct {
	ID        uuid.UUID `json:"id" bson:"_id"`
	Nickname  string    `json:"nickname" bson:"nickname"`
	Name      string    `json:"name" bson:"nickname"`
	Birthdate string    `json:"birthdate" bson:"birthdate"`
	Stack     []string  `json:"stack" bson:"stack"`
}
