package main

import (
	"fiberGorm/route"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// Initial Route
	route.RouteInit(app)

	app.Listen(":3000")
}
