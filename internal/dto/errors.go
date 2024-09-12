package dto

import (
	"github.com/gofiber/fiber/v3"
)

const (
	FieldRequired        = "FIELD_REQUIRED"
	FieldBadType         = "FIELD_BADTYPE"
	FieldBadFormat       = "FIELD_BADFORMAT"
	FieldIncorrect       = "FIELD_INCORRECT"
	ServiceUnavailable   = "SERVICE_UNAVAILABLE"
	OrderNotActivate     = "ORDER_NOT_ACTIVATED"
	FieldRequiredMsg     = "is required"
	FieldBadTypeMsg      = "has incorrect type"
	FieldBadFormatMsg    = "has incorrect format"
	FieldIncorrectMsg    = "is incorrect"
	InternalError        = "Service is currently unavailable. Please try again later."
	OrderNotActivateDesc = "The order has not been activated yet"
)

type Response struct {
	Status string `json:"status"`
	Error  *Error `json:"error,omitempty"`
	Data   any    `json:"data,omitempty"`
}

type Error struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

func BadResponseError(ctx fiber.Ctx, code, desc string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(Response{
		Status: "error",
		Error: &Error{
			Code: code,
			Desc: desc,
		},
	})
}

func InternalServerError(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Status: "error",
		Error: &Error{
			Code: ServiceUnavailable,
			Desc: InternalError,
		},
	})
}
