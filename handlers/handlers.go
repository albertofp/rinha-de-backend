package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/albertofp/rinha-de-backend/database"
)

func PostPerson(c *fiber.Ctx) error {
	return nil
}

func SearchPerson(c *fiber.Ctx, query string) error {
	return nil
}

func CountPeople(c *fiber.Ctx) error {

	return nil
}

func GetPersonByID(c *fiber.Ctx, id string) error {

	return nil
}

func TestHandler(c *fiber.Ctx) error {
	testDoc := bson.M{"name": "testDoc"}
	collection := database.GetCollection("pessoas")
	newDoc, err := collection.InsertOne(context.TODO(), testDoc)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error inserting doc")
	}

	return c.JSON(newDoc)
}
