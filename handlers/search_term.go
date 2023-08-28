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

// Query godoc
// @Summary Search by term
// @Description Search for a person in database by a given query string. Search term must not be empty
// @Tags pessoas
// @Param t query string false "Search term"
// @Produce json
// @Success 200 {array} models.PersonDTO{}
// @Failure 400 {object} models.ErrorResponse{}
// @Failure 500 {object} models.ErrorResponse{}
// @Router /pessoas [get]
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
			{"id": bson.M{"$regex": t.T, "$options": "i"}},
			{"apelido": bson.M{"$regex": t.T, "$options": "i"}},
			{"nome": bson.M{"$regex": t.T, "$options": "i"}},
			{"stack": bson.M{"$regex": t.T, "$options": "i"}},
		},
	}

	opts := options.Find().SetLimit(50)

	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		return err
	}
	defer cursor.Close(context.TODO())

	var results []models.PersonDTO
	for cursor.Next(context.TODO()) {
		var person models.PersonDTO
		err := cursor.Decode(&person)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		results = append(results, person)
	}
	fmt.Println("results: ", &results)
	if results == nil {
		emptyArr := make([]models.PersonDTO, 0)
		return c.Status(fiber.StatusOK).JSON(emptyArr)
	}

	return c.Status(fiber.StatusOK).JSON(results)
}
