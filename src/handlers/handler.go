package handlers

import (
	"mini-bank/src/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Handler struct {
	App       *fiber.App
	DB        *gorm.DB
	Contract  *services.Contract
	Validator *validator.Validate
}
