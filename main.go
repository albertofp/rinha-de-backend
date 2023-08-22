package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
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

	app.Get("/contagem-pessoas", handlers.CountPeople)
	app.Get("/healthcheck", handlers.Healthcheck)
	app.Get("/getall", handlers.GetAllPerson)
	//app.Get("pessoas/:id", handlers.GetPersonById)
	app.Post("/pessoas", handlers.PostPerson)

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
