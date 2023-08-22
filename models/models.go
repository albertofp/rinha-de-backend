package models

import "github.com/google/uuid"

type PersonDTO struct {
	Id         uuid.UUID `json:"id"         bson:"id"`
	Apelido    string    `json:"apelido"    bson:"apelido"`
	Nome       string    `json:"nome"       bson:"nome"`
	Nascimento string    `json:"nascimento" bson:"nascimento"`
	Stack      []string  `json:"stack"      bson:"stack"`
}
