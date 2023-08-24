package models

type PersonDTO struct {
	Id         string   `json:"id"         bson:"id"`
	Apelido    string   `json:"apelido"    bson:"apelido"`
	Nome       string   `json:"nome"       bson:"nome"`
	Nascimento string   `json:"nascimento" bson:"nascimento"`
	Stack      []string `json:"stack"      bson:"stack"`
}
type PersonCreateRequest struct {
	Apelido    string   `json:"apelido"    bson:"apelido"`
	Nome       string   `json:"nome"       bson:"nome"`
	Nascimento string   `json:"nascimento" bson:"nascimento"`
	Stack      []string `json:"stack"      bson:"stack"`
}

type PersonCreateResponse struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
