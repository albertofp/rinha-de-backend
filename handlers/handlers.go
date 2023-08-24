package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/models"
)

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

func Status(c *fiber.Ctx) error {
	return c.SendString("GET Success")
}
