package poc

import (
	poc_model "poc_demo/src/components/poc/model"
	validator "poc_demo/src/components/validation"

	"github.com/gofiber/fiber/v2"
)

func PocPost(c *fiber.Ctx) error {

	user := new(poc_model.PocModel)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})

	}

	errors := validator.ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	return c.JSON(user)
}
