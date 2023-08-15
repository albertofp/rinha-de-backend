package main

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/albertofp/rinha-de-backend/database"
)

type Person struct {
	ID        uuid.UUID
	Nickname  string
	Name      string
	Birthdate string
	Stack     []string
}

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}

	defer database.CloseDB()

	app := fiber.New()

	app.Get("/pessoas", func(c *fiber.Ctx) error {
		return c.SendString("GET /pessoas")
	})

	app.Post("/pessoas", func(c *fiber.Ctx) error {
		testDoc := bson.M{"name": "testDoc"}
		collection := database.GetCollection("pessoas")
		newDoc, err := collection.InsertOne(context.TODO(), testDoc)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error inserting doc")
		}

		return c.JSON(newDoc)
	})

	app.Listen(":8080")
}

func initApp() error {
	err := loadENV()
	if err != nil {
		return err
	}

	err = database.InitDB()
	if err != nil {
		return err
	}
	return nil
}

func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
