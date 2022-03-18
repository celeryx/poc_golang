package main

import (
	pocController "poc_demo/src/components/poc/controller"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/v1/poc", pocController.PocPost)
}

func main() {

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
