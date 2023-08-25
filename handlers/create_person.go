package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/models"
	"github.com/albertofp/rinha-de-backend/validation"
)

//@BasePath /pessoas

//@Summary Create new person document
//@Accept json
//@Produce json
//@Param request body PersonCreateRequest true "Request body"
//@Success 201 {object} PersonCreateResponse
//@Failure 422 {object} ErrorResponse
//@Failure 500 {object} ErrorResponse
//@Router /pessoas [post]

func Create(c *fiber.Ctx) error {
	//TODO: check if nickname already in db -> skip
	newPerson := new(models.PersonDTO)
	if err := c.BodyParser(newPerson); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
	}

	statusCode, res := validation.ValidateRequest(newPerson)
	if statusCode != fiber.StatusCreated {
		return c.Status(statusCode).JSON(fiber.Map{"error": res})
	}

	if personExists, err := checkPersonExists(newPerson.Apelido); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	} else if personExists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Person with the same 'apelido' already exists"})
	}

	newPerson.Id = uuid.New().String()
	//filter := bson.M{"apelido": newPerson.Apelido}

	coll := database.GetCollection("pessoas")
	_, err := coll.InsertOne(context.TODO(), newPerson)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Location", fmt.Sprintf("/pessoas/%s", newPerson.Id))
	return c.Status(statusCode).JSON(fiber.Map{"success": res, "id": newPerson.Id})
}

func checkPersonExists(apelido string) (bool, error) {
	coll := database.GetCollection("pessoas")

	filter := bson.M{"apelido": apelido}
	count, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
