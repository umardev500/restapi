package handler

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/restapi/domain"
	"github.com/umardev500/restapi/domain/model"
	"github.com/umardev500/restapi/helper"
)

type productHandler struct {
	service   domain.ProductService
	validator *validator.Validate
}

func NewProductHandler(service domain.ProductService, validator *validator.Validate) domain.ProductHanlder {
	return &productHandler{
		service:   service,
		validator: validator,
	}
}

func (ph *productHandler) Create(c *fiber.Ctx) error {
	var product *model.ProductStore = new(model.ProductStore)
	if err := c.BodyParser(product); err != nil {
		return helper.ErrResponse(c, fiber.StatusUnprocessableEntity, nil, err.Error())
	}

	// validate the struct
	if err := ph.validator.Struct(product); err != nil {
		var details []string
		for _, fieldErr := range err.(validator.ValidationErrors) {
			details = append(details, fieldErr.Error())
		}

		return helper.ErrResponse(c, fiber.StatusUnprocessableEntity, details)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := ph.service.CreateProduct(ctx, *product)
	if err != nil {
		return helper.ErrResponse(c, fiber.StatusInternalServerError, nil, err.Error())
	}

	return helper.OkResponse(c, fiber.StatusOK, "Product created successfully", nil)
}
