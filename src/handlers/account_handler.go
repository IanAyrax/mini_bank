package handlers

import (
	"mini-bank/src/requests"
	"mini-bank/src/responses"
	"mini-bank/src/services"
	"mini-bank/src/view_models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	Handler
}

func NewAccountHandler(handler Handler) *AccountHandler {
	return &AccountHandler{Handler: handler}
}

func (handler AccountHandler) Create(ctx *fiber.Ctx) (err error) {
	req := new(requests.AccountRequest)
	if err := ctx.BodyParser(req); err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	//Validate
	if err := handler.Validator.Struct(req); err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	accountService := services.NewAccountService(handler.Contract)
	data, err := accountService.Create(req)
	if err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	return responses.ReponseSuccess(ctx, data, "Success to create account")
}

func (handler AccountHandler) List(ctx *fiber.Ctx) (err error) {
	req := new(requests.FilterQueryAccountRequest)

	if err := ctx.QueryParser(req); err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	accountService := services.NewAccountService(handler.Contract)
	data, err := accountService.List(req)

	if err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	detail := false
	if req != nil {
		if req.Detail == true {
			detail = req.Detail
		}
	}

	if detail == false {
		res := view_models.NewAccountVm().BuildListNoFromVmList(data)

		return responses.ReponseSuccess(ctx, res, "Success to list account")
	}

	return responses.ReponseSuccess(ctx, data, "Success to list account")
}

func (handler AccountHandler) Detail(ctx *fiber.Ctx) (err error) {
	accountID := ctx.Params("id")

	accountService := services.NewAccountService(handler.Contract)
	data, err := accountService.Detail(accountID)

	if err != nil {
		return responses.ReponseError(ctx, http.StatusBadRequest, err)
	}

	if data == nil {
		return responses.ReponseSuccess(ctx, data, "Data doesn't exist")
	}

	return responses.ReponseSuccess(ctx, data, "Success to get account")
}

func (handler AccountHandler) Update(ctx *fiber.Ctx) (err error) {
	return ctx.Status(http.StatusBadRequest).SendString("Not Implemented Yet")
}

func (handler AccountHandler) Delete(ctx *fiber.Ctx) (err error) {
	return ctx.Status(http.StatusBadRequest).SendString("Not Implemented Yet")
}
