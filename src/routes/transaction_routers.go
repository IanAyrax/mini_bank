package routes

import (
	"mini-bank/src/handlers"

	"github.com/gofiber/fiber/v2"
)

type TransactionRouter struct {
	RouteGroup fiber.Router
	Handler    handlers.Handler
}

func (route TransactionRouter) TransactionRouter() {
	transactionHandler := handlers.NewTransactionHandler(route.Handler)

	transactionRoutes := route.RouteGroup.Group("/transaction")
	transactionRoutes.Get("/list", transactionHandler.List)
	transactionRoutes.Get("", transactionHandler.Detail)
	transactionRoutes.Post("", transactionHandler.Create)
	transactionRoutes.Post("/search", transactionHandler.Search)
	transactionRoutes.Patch("/:id", transactionHandler.Update)
	transactionRoutes.Delete("/:id", transactionHandler.Delete)
}
