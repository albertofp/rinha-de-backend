package models

import "github.com/google/uuid"

type CreatePerson struct {
	Nickname  string   `json:"apelido" bson:"apelido" validate:"required,max=32"`
	Name      string   `json:"nome" bson:"nome" validate:"required,max=100"`
	Birthdate string   `json:"nascimento" bson:"nascimento" validate:"required,2023-08-15"`
	Stack     []string `json:"stack" bson:"stack" validate:"dive,max=32"`
}

type PersonResponse struct {
	ID        uuid.UUID `json:"id"`
	Nickname  string    `json:"apelido"`
	Name      string    `json:"nome"`
	Birthdate string    `json:"nascimento"`
	Stack     []string  `json:"stack"`
}
