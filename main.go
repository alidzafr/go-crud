package main

import (
	"fiberGorm/controllers/bookcontroller"
	"fiberGorm/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	// Initial Route (dinonaktifkan karena agak ribet)
	// Rename handler into controllers
	// route.RouteInit(app)

	apiBook := app.Group("/api/books")

	apiBook.Get("/", bookcontroller.Index)
	apiBook.Get("/:id", bookcontroller.Show)
	apiBook.Post("/", bookcontroller.Create)
	apiBook.Put("/:id", bookcontroller.Update)
	apiBook.Delete("/:id", bookcontroller.Delete)

	app.Listen(":8000")
}
