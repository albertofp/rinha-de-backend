package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/models"
	"github.com/albertofp/rinha-de-backend/validation"
)

func Create(c *fiber.Ctx) error {
	//TODO: check if nickname already in db -> skip
	newPerson := new(models.PersonDTO)
	if err := c.BodyParser(newPerson); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{})
	}

	statusCode, res := validation.ValidateRequest(newPerson)
	if statusCode != fiber.StatusCreated {
		return c.Status(statusCode).JSON(fiber.Map{"error": res})
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
