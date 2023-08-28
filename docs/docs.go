// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Alberto F. Pluecker",
            "url": "https://github.com/albertofp",
            "email": "albertopluecker@gmail.com"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/contagem-pessoas": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "count"
                ],
                "summary": "Count total amount of people in the database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CountResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/getall": {
            "get": {
                "description": "Returns an empty array if no people found.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "getall"
                ],
                "summary": "Get every person in the database",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PersonDTO"
                            }
                        }
                    }
                }
            }
        },
        "/pessoas": {
            "get": {
                "description": "Search for a person in database by a given query string. Search term must not be empty",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pessoas"
                ],
                "summary": "Search by term",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search term",
                        "name": "t",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PersonDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a person to the database.  Returns an error if another person with the same value for the \"apelido\" field exists. Apelido and Nome have to be strings of length up to 32 and 100, respectively.  Nascimento has to follow date format YYYY-MM-DD. Stack is optional, but each entry contained has to be a string of up to 32 chars in length.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pessoas"
                ],
                "summary": "Create new person document",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PersonCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.PersonCreateResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/pessoas/{id}": {
            "get": {
                "description": "Returns a person with the given id (UUID format)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pessoas"
                ],
                "summary": "Search person by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Person ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PersonDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/status": {
            "get": {
                "tags": [
                    "status"
                ],
                "summary": "Health check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CountResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.PersonCreateRequest": {
            "type": "object",
            "properties": {
                "apelido": {
                    "type": "string"
                },
                "nascimento": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "stack": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.PersonCreateResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.PersonDTO": {
            "type": "object",
            "properties": {
                "apelido": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nascimento": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "stack": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Rinha de Backend Q3 2023 - Alberto Pluecker",
	Description:      "Docs auto-generated by Swagger",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
