package helper

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/umardev500/restapi/domain/model"
)

func OkResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	payload := model.OkResponse{
		Code:    code,
		Success: true,
		Message: message,
		Data:    data,
	}
	return c.Status(code).JSON(payload)
}

func ErrResponse(c *fiber.Ctx, code int, details []string, messages ...string) error {
	payload := model.ErrResponse{
		Code:    code,
		Success: false,
		Details: details,
		Message: utils.StatusMessage(code),
	}
	if len(messages) > 0 {
		payload.Message = messages[0]
	}

	return c.Status(code).JSON(payload)
}
