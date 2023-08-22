package handlers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/models"
	"github.com/albertofp/rinha-de-backend/validation"
)

func Healthcheck(c *fiber.Ctx) error {
	return c.SendString("GET Success")
}

func SearchPerson(c *fiber.Ctx) error {
	t := c.Params("t")
	fmt.Printf("t = %s\n", t)
	if t == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "400 - Bad request: query must not be empty",
		})
	}
	coll := database.GetCollection("pessoas")
	model := mongo.IndexModel{Keys: bson.D{{"description", "text"}}}
	name, err := coll.Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		panic(err)
	}
	filter := bson.D{{"$text", bson.D{{"$search", t}}}}
	fmt.Println("Name of index created: " + name)

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []models.PersonDTO
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}

	return c.Status(fiber.StatusOK).JSON(results)
}

func PostPerson(c *fiber.Ctx) error {
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
	fmt.Printf("id: %s", id)
	coll := database.GetCollection("pessoas")
	filter := bson.D{primitive.E{Key: "id", Value: id}}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	res, _ := json.Marshal(cursor)
	return c.Status(fiber.StatusFound).JSON(res)
}

func GetAllPerson(c *fiber.Ctx) error {
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
	return c.JSON(people)

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
