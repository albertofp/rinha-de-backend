package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/albertofp/rinha-de-backend/database"
	"github.com/albertofp/rinha-de-backend/models"
)

func Healthcheck(c *fiber.Ctx) error {
	return c.SendString("GET Success")
}

func PostPerson(c *fiber.Ctx) error {
	newPerson := new(models.PersonDTO)
	if err := c.BodyParser(newPerson); err != nil {
		return err
	}
	newPerson.Id = uuid.New()
	coll := database.GetCollection("pessoas")
	nDoc, err := coll.InsertOne(context.TODO(), newPerson)
	if err != nil {
		return err
	}

	statusCode, msg := validateRequest(newPerson)
	if statusCode != fiber.StatusCreated {
		return c.Status(statusCode).JSON(fiber.Map{"error": msg})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": msg, "id": nDoc.InsertedID})
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

func GetPersonById(c *fiber.Ctx, id string) error {

	return nil
}

func validateRequest(req *models.PersonDTO) (int, string) {

	msgUnprocessable := "422 - Unprocessable entity"
	msgBadRequest := "401 - Bad request"
	msgCreated := "201 - Created"

	if len(req.Apelido) > 32 || len(req.Nome) > 100 {
		return fiber.StatusUnprocessableEntity, msgUnprocessable
	}

	if !validateDate(req) {
		return fiber.StatusUnprocessableEntity, msgUnprocessable
	}

	if !isString(req.Apelido) || !isString(req.Nome) {
		return fiber.StatusBadRequest, msgBadRequest
	}

	if req.Apelido == "" || req.Nome == "" {
		return fiber.StatusUnprocessableEntity, msgUnprocessable
	}

	if req.Stack != nil {
		for i := range req.Stack {
			if len(req.Stack[i]) > 32 || !isString(req.Stack[i]) {
				fmt.Println("stack not string")
				return fiber.StatusUnprocessableEntity, msgUnprocessable
			}
		}

	}
	if !validateDate(req) {
		return fiber.StatusUnprocessableEntity, msgUnprocessable
	}

	return fiber.StatusCreated, msgCreated
}

func validateDate(req *models.PersonDTO) bool {
	dateLayout := "2002-12-17"
	if _, err := time.Parse(dateLayout, req.Nascimento); err != nil {
		return true
	}
	return false
}
func isString[T any](input T) bool {
	var check interface{}
	check = input
	_, ok := check.(string)
	return ok
}
