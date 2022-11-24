package routes

import (
	"mini-bank/src/handlers"

	"github.com/gofiber/fiber/v2"
)

type AccountRouter struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route AccountRouter) AccountRouter() {
	accountHandler := handlers.NewAccountHandler(route.Handler)

	accountRoutes := route.RouteGroup.Group("/account")
	accountRoutes.Get("", accountHandler.List)
	accountRoutes.Get("/:id", accountHandler.Detail)
	accountRoutes.Post("", accountHandler.Create)
	accountRoutes.Patch("/:id", accountHandler.Update)
	accountRoutes.Delete("/:id", accountHandler.Delete)
}
