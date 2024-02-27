package main

import (
	"fiberGorm/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	// Initial Route (dinonaktifkan karena agak ribet)
	// Rename handler into controllers
	// route.RouteInit(app)

	app.Listen(":8000")
}
