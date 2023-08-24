package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/models"
)

func SearchId(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	coll := database.GetCollection("pessoas")
	filter := bson.M{"id": id}

	var pessoa models.PersonDTO
	err := coll.FindOne(context.TODO(), filter).Decode(&pessoa)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(pessoa)
}
