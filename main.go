package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/handlers"
)

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}

	defer database.CloseDB()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(requestid.New())

	p := app.Group("/pessoas")

	app.Get("/pessoas", handlers.GetPersonByTerm)
	p.Get("/:id", handlers.GetPersonById)
	p.Post("/", handlers.PostPerson)

	app.Get("/contagem-pessoas", handlers.Count)
	app.Get("/getall", handlers.GetAll)
	app.Get("/status", handlers.Status)

	port := os.Getenv("PORT")
	app.Listen(":" + port)
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
