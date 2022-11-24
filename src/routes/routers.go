package routes

import (
	"mini-bank/src/handlers"
	"mini-bank/src/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Routers struct {
	App       *fiber.App
	DB        *gorm.DB
	Validator *validator.Validate
}

func (routers Routers) Routers() {
	rootAccountGroup := routers.App.Group("/mini-bank")

	contract := services.Contract{
		App: routers.App,
		DB:  routers.DB,
	}

	handler := handlers.Handler{
		App:       routers.App,
		DB:        routers.DB,
		Contract:  &contract,
		Validator: routers.Validator,
	}
	accountRouter := AccountRouter{
		RouteGroup: rootAccountGroup,
		Handler:    handler,
	}
	accountRouter.AccountRouter()

	transactionRouter := TransactionRouter{
		RouteGroup: rootAccountGroup,
		Handler:    handler,
	}
	transactionRouter.TransactionRouter()
}
