package handlers

import (
	"context"
	"fmt"

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
	coll := database.GetCollection("pessoas")
	filter := bson.D{}
	count, err := coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error counting people")
	}
	estCount, err := coll.EstimatedDocumentCount(context.TODO())
	if err != nil {
		return err
	}

	fmt.Printf("Estimated Count: %d\n", estCount)
	fmt.Printf("Accurate Count: %d\n", count)
	return nil
}

func GetPersonByID(c *fiber.Ctx, id string) error {

	return nil
}

func TestHandler(c *fiber.Ctx) error {
	testDoc := bson.M{"name": "testDoc"}
	newDoc, err := database.GetCollection("pessoas").InsertOne(context.TODO(), testDoc)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error inserting doc")
	}

	return c.JSON(newDoc)
}
