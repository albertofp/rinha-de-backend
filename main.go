package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	"github.com/albertofp/rinha-de-backend/database"
	_ "github.com/albertofp/rinha-de-backend/docs"
	"github.com/albertofp/rinha-de-backend/handlers"
)

//@title Rinha de Backend Q3 2023 - Alberto Pluecker
//@version 1.0
//@license.name MIT
//@contact.name Alberto F. Pluecker
//@contact.url https://github.com/albertofp
//@contact.email albertopluecker@gmail.com
//@host      localhost:8080

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}
	defer database.CloseDB()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cache.New())
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	app.Get("/pessoas", handlers.Query)
	app.Get("/pessoas/:id", handlers.SearchId)
	app.Post("/pessoas", handlers.Create)

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
