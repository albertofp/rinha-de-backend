package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/models"
	"github.com/albertofp/rinha-de-backend/validation"
)

type QueryParams struct {
	T string `query:"t"`
}

func GetPersonByTerm(c *fiber.Ctx) error {
	t := new(QueryParams)
	if err := c.QueryParser(t); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	fmt.Println("t.T: ", t.T)

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

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return err
	}
	var results []models.PersonDTO
	if err = cursor.All(context.TODO(), &results); err != nil {
		return err
	}

	fmt.Println("results: ", &results)

	return c.Status(fiber.StatusOK).JSON(results)
}

func PostPerson(c *fiber.Ctx) error {
	//TODO: check if nickname already in db -> skip
	newPerson := new(models.PersonDTO)
	if err := c.BodyParser(newPerson); err != nil {
		return err
	}
	newPerson.Id = uuid.New()
	coll := database.GetCollection("pessoas")
	_, err := coll.InsertOne(context.TODO(), newPerson)
	if err != nil {
		return err
	}

	statusCode, msg := validation.ValidateRequest(newPerson)
	if statusCode != fiber.StatusCreated {
		return c.Status(statusCode).JSON(fiber.Map{"error": msg})
	}

	header := fmt.Sprintf("/pessoas/%s", newPerson.Id)
	c.Set("Location", header)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": msg, "id": newPerson.Id})
}

func GetPersonById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}
	fmt.Printf("id: %s", id)
	coll := database.GetCollection("pessoas")
	filter := bson.D{primitive.E{Key: "id", Value: id}}

	pessoa := models.PersonDTO{}
	err := coll.FindOne(context.TODO(), filter).Decode(&pessoa)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusFound).JSON(pessoa)
}

func GetAll(c *fiber.Ctx) error {
	filter := bson.M{}
	coll := database.GetCollection("pessoas")

	var people []models.PersonDTO
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return err
	}
	if err = cursor.All(context.TODO(), &people); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(people)

}

func Count(c *fiber.Ctx) error {
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

func Status(c *fiber.Ctx) error {
	return c.SendString("GET Success")
}
