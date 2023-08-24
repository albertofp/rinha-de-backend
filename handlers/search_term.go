package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/models"
)

type QueryParams struct {
	T string //`query:"t"`
}

func Query(c *fiber.Ctx) error {
	t := new(QueryParams)
	if err := c.QueryParser(t); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	if t.T == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "400 - Bad request: query must not be empty",
		})
	}
	coll := database.GetCollection("pessoas")
	filter := bson.M{
		"$or": []bson.M{
			{"apelido": t.T},
			{"nome": t.T},
			{"stack": t.T},
		},
	}
	opts := options.Find().SetLimit(50)

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		return err
	}
	var results []models.PersonDTO
	if err = cursor.All(context.TODO(), &results); err != nil {
		return err
	}

	fmt.Println("results: ", &results)
	if results == nil {
		emptyArr := make([]models.PersonDTO, 0)
		return c.Status(fiber.StatusOK).JSON(emptyArr)
	}

	return c.Status(fiber.StatusOK).JSON(results)
}
