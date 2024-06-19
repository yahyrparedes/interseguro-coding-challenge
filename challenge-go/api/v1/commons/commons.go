package commons

import (
	"github.com/gofiber/fiber/v2"
	"reflect"
	"strings"
)

func FieldJSONName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

	if name == "-" {
		return ""
	}

	return name
}

type (
	ErrorResponse struct {
		Message string      `json:"message"`
		Status  int         `json:"status"`
		Data    interface{} `json:"data,omitempty"`
	}
	SuccessResponse struct {
		Message string      `json:"message"`
		Status  int         `json:"status"`
		Data    interface{} `json:"data,omitempty"`
	}
	IError struct {
		Field string `json:"field"`
		Tag   string `json:"tag"`
		Value string `json:"value,omitempty" `
	}
)

func BadRequestAndData(c *fiber.Ctx, message string, status int, data interface{}) error {
	return c.Status(fiber.StatusBadRequest).
		JSON(ErrorResponse{Message: message, Status: status, Data: data})

}

func BadRequestAndMessage(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).
		JSON(ErrorResponse{Message: message, Status: fiber.StatusBadRequest})
}

func BadRequest(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusBadRequest)
}

func GenericErrorResponse(ctx *fiber.Ctx, err ErrorResponse) error {
	return ctx.Status(err.Status).JSON(err)
}

func OkAndData(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{Message: "Success", Status: fiber.StatusOK, Data: data})
}
