package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tigerbig/spatial-data-plateform/internal/domain/collection"
	"github.com/tigerbig/spatial-data-plateform/internal/usecase"
)

type SpatialHandlers struct {
	useCase *usecase.SpatialUseCases
}

func NewSpatialHandlers(useCase *usecase.SpatialUseCases) *SpatialHandlers {
	return &SpatialHandlers{
		useCase: useCase,
	}
}

func (handler *SpatialHandlers) Create(c fiber.Ctx) error {
	var req collection.Features

	errParse := c.Bind().Body(&req)
	if errParse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": errParse.Error()})
	}
	data, err := handler.useCase.Create(c.Context(), &req)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(data)
}

func (handler *SpatialHandlers) GetAll(c fiber.Ctx) error {
	data, err := handler.useCase.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": data,
	})
}
