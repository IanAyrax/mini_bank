package handlers

import (
	"errors"
	"mini-bank/src/requests"
	"mini-bank/src/responses"
	"mini-bank/src/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	Handler
}

func NewTransactionHandler(handler Handler) *TransactionHandler {
	return &TransactionHandler{Handler: handler}
}

func (handler TransactionHandler) Create(ctx *fiber.Ctx) (err error) {
	req := new(requests.TransactionRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	//Validate
	if err := handler.Validator.Struct(req); err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	transactionService := services.NewTransactionService(handler.Contract)
	data, err := transactionService.Create(req)
	if err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	return responses.ReponseSuccess(ctx, data, "Success to create transaction")
}

func (handler TransactionHandler) Search(ctx *fiber.Ctx) (err error) {
	req := new(requests.FilterTransactionRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	transactionService := services.NewTransactionService(handler.Contract)
	data, err := transactionService.List(req)
	if err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	return responses.ReponseSuccess(ctx, data, "Success to search transaction")
}

func (handler TransactionHandler) List(ctx *fiber.Ctx) (err error) {
	transactionService := services.NewTransactionService(handler.Contract)
	data, err := transactionService.List(&requests.FilterTransactionRequest{})

	if err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	return responses.ReponseSuccess(ctx, data, "Success to list transaction")
}

func (handler TransactionHandler) Detail(ctx *fiber.Ctx) (err error) {
	transactionIDByte := ctx.Request().Header.Peek("X-User-Id")

	if transactionIDByte == nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, errors.New("Unauthorized"))
	}

	transactionID := string(transactionIDByte)

	if &transactionID == nil || transactionID == "" {
		return responses.ReponseError(ctx, http.StatusBadRequest, errors.New("Unauthorized"))
	}

	//Check if ID is uuid
	_, err = uuid.Parse(transactionID)
	if err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	transactionService := services.NewTransactionService(handler.Contract)
	data, err := transactionService.Detail(transactionID)

	if err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	if data == nil {
		return responses.ReponseSuccess(ctx, data, "Data doesn't exist")
	}

	return responses.ReponseSuccess(ctx, data, "Success to get transaction")
}

func (handler TransactionHandler) Update(ctx *fiber.Ctx) (err error) {
	return ctx.Status(http.StatusBadRequest).SendString("Not Implemented Yet")
}

func (handler TransactionHandler) Delete(ctx *fiber.Ctx) (err error) {
	return ctx.Status(http.StatusBadRequest).SendString("Not Implemented Yet")
}
