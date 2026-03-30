package route

import (
	"github.com/gofiber/fiber/v3"
	"github.com/tigerbig/spatial-data-plateform/internal/delivery/handler"
)

// func SetRoute(app *fiber.App, spatialHandler *handler.SpatialHandler) {
func SetRoute(app *fiber.App, spatialHandler *handler.SpatialHandlers) {
	api := app.Group("/api")
	api.Post("/create-location", spatialHandler.Create)
	api.Get("/get-list-spatial", spatialHandler.GetAll)
}
