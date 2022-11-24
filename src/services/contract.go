package services

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Contract struct {
	App *fiber.App
	DB  *gorm.DB
}
