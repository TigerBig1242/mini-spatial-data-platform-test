package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v3"

	"github.com/tigerbig/spatial-data-plateform/internal/config"
	"github.com/tigerbig/spatial-data-plateform/internal/delivery/handler"
	router "github.com/tigerbig/spatial-data-plateform/internal/delivery/route"
	"github.com/tigerbig/spatial-data-plateform/internal/infrastructure/database"
	infraRepo "github.com/tigerbig/spatial-data-plateform/internal/infrastructure/repository"
	"github.com/tigerbig/spatial-data-plateform/internal/usecase"
)

func main() {
	runMigrate := flag.Bool("migrate", false, "Run database migration")
	flag.Parse()

	databaseConfig := config.LoadConfig()

	client, err := database.ConnectDatabase(databaseConfig)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("spatial-data-test")

	spatialRepo := infraRepo.NewSpatialRepo(db)

	spatialUseCase := usecase.NewSpatialUseCases(spatialRepo)

	spatialHandler := handler.NewSpatialHandlers(spatialUseCase)

	if *runMigrate {
		migrateErr := database.Migrate(db)
		if migrateErr != nil {
			panic(migrateErr)
		}
	}

	app := fiber.New()

	app.Get("/spatial", func(c fiber.Ctx) error {
		return c.SendString("Spatial Data")
	})

	router.SetRoute(app, spatialHandler)

	app.Listen(":8080")
}
