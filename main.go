package main

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

	"github.com/albertofp/rinha-de-backend/database"
	_ "github.com/albertofp/rinha-de-backend/docs"
	"github.com/albertofp/rinha-de-backend/handlers"
)

//@title Rinha de Backend Q3 2023 - Alberto Pluecker
//@version 1.0
//@description Docs auto-generated by Swagger
//@termsOfService http://swagger.io/terms/

//@contact.name Alberto F. Pluecker
//@contact.url https://github.com/albertofp
//@contact.email albertopluecker@gmail.com

//@license.name MIT

//@host localhost:8080
//@BasePath /

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
	//r.Use(cache.New())
	r.Get("/docs/*", swagger.HandlerDefault)

	r.Get("/pessoas", handlers.Query)
	r.Get("/pessoas/:id", handlers.SearchId)
	r.Post("/pessoas", handlers.Create)

	r.Get("/contagem-pessoas", handlers.Count)
	r.Get("/getall", handlers.GetAll)
	r.Get("/status", handlers.Status)

	r.Listen(":8080")
}

func initApp() error {
	err := database.InitDB()
	if err != nil {
		return err
	}
	return nil
}
