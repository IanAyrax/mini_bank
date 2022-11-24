package responses

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FactoryBaseResponse struct {
	Data interface{} `json:"data"`
	// Meta   interface{}           `json:"meta"`
	Status FactoryStatusResponse `json:"status"`
}

type FactoryStatusResponse struct {
	Success      bool        `json:"success"`
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	ErrorMessage interface{} `json:"error_message"`
}

func ReponseSuccess(ctx *fiber.Ctx, data interface{}, message string) error {
	resp := &FactoryBaseResponse{
		Data: data,
		// Meta: nil,
		Status: FactoryStatusResponse{
			Success:      true,
			Code:         http.StatusOK,
			Message:      message,
			ErrorMessage: nil,
		},
	}

	return ctx.Status(http.StatusOK).JSON(resp)
}

func ReponseError(ctx *fiber.Ctx, errorCode int, err error) error {
	resp := &FactoryBaseResponse{
		Data: nil,
		// Meta: nil,
		Status: FactoryStatusResponse{
			Success:      false,
			Code:         errorCode,
			Message:      "Error",
			ErrorMessage: err.Error(),
		},
	}

	return ctx.Status(errorCode).JSON(resp)
}
