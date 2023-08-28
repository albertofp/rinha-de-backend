package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/models"
)

// GetAll godoc
// @Summary Get every person in the database
// @Description Returns an empty array if no people found.
// @Tags getall
// @Produce json
// @Success 200 {array} models.PersonDTO
// @Router /getall [get]
func GetAll(c *fiber.Ctx) error {
	filter := bson.M{}
	coll := database.GetCollection("pessoas")

	var people []models.PersonDTO
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {

	}
	if err = cursor.All(context.TODO(), &people); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(people)

}

// Status godoc
// @Summary Health check
// @Tags status
// @Success 200 {array} string
// @Router /status [get]
func Status(c *fiber.Ctx) error {
	return c.SendString("GET Success")
}
