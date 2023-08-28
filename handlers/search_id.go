package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/models"
)

// SearchId godoc
// @Summary Search person by ID
// @Description Returns a person with the given id (UUID format)
// @Tags pessoas
// @Param id path string true "Person ID"
// @Produce json
// @Success 200 {object} models.PersonDTO{}
// @Failure 400 {object} models.ErrorResponse{}
// @Failure 404 {object} models.ErrorResponse{}
// @Failure 500 {object} models.ErrorResponse{}
// @Router /pessoas/{id} [get]
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
