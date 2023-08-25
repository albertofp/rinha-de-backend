package main

import (
	"os"

	"github.com/bytedance/sonic"
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

	r := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	r.Use(logger.New())
	r.Use(cache.New())
	r.Get("/swagger/*", fiberSwagger.WrapHandler)

	r.Get("/pessoas", handlers.Query)
	r.Get("/pessoas/:id", handlers.SearchId)
	r.Post("/pessoas", handlers.Create)

	r.Get("/contagem-pessoas", handlers.Count)
	r.Get("/getall", handlers.GetAll)
	r.Get("/status", handlers.Status)

	r.Listen(":" + os.Getenv("PORT"))
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
