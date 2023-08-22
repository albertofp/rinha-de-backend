package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID        string
	Nickname  string
	Name      string
	Birthdate string
	Stack     []string
}

type CreatePerson struct {
	Nickname  string   `json:"apelido" bson:"apelido"`
	Name      string   `json:"nome" bson:"nome"`
	Birthdate string   `json:"nascimento" bson:"nascimento"`
	Stack     []string `json:"stack" bson:"stack"`
}

type PersonResponse struct {
	ID        uuid.UUID `json:"id"`
	Nickname  string    `json:"apelido"`
	Name      string    `json:"nome"`
	Birthdate string    `json:"nascimento"`
	Stack     []string  `json:"stack"`
}

func (c *CreatePerson) Validate() error {
	var invalidReqError = errors.New("invalid request")
	dateFormat := "2000-01-01"

	if len(c.Name) > 100 {
		return invalidReqError
	}

	if len(c.Nickname) > 32 {
		return invalidReqError
	}

	if _, err := time.Parse(dateFormat, c.Birthdate); err != nil {
		return invalidReqError
	}

	for _, skill := range c.Stack {
		if len(skill) > 32 {
			return invalidReqError
		}
	}

	return nil
}
