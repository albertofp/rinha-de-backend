package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/models"
)

func Healthcheck(c *fiber.Ctx) error {
	return c.SendString("GET Success")
}

func PostPerson(c *fiber.Ctx) error {
	newPerson := new(models.PersonDTO)
	if err := c.BodyParser(newPerson); err != nil {
		return err
	}
	coll := database.GetCollection("pessoas")
	nDoc, err := coll.InsertOne(context.TODO(), newPerson)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"id": nDoc.InsertedID,
	})
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

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"count": count,
	})
}

func GetPersonById(c *fiber.Ctx, id string) error {

	return nil
}
