package models

type PersonDTO struct {
	Id         string   `json:"id"         bson:"id"`
	Apelido    string   `json:"apelido"    bson:"apelido"`
	Nome       string   `json:"nome"       bson:"nome"`
	Nascimento string   `json:"nascimento" bson:"nascimento"`
	Stack      []string `json:"stack"      bson:"stack"`
}
