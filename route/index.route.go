package route

import (
	"fiberGorm/handler"

	"github.com/gofiber/fiber/v3"
)

func RouteInit(r *fiber.App) {
	r.Get("/", handler.UserHandlerRead)
}
