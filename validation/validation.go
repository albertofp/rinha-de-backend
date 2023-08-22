package validation

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/albertofp/rinha-de-backend/models"
)

func ValidateRequest(req *models.PersonDTO) (int, string) {

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
