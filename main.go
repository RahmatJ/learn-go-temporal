package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"learn-temporal-go/internal/route"
	"learn-temporal-go/internal/workflow"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msgf("Starting Server")
	app := fiber.New()
	port := 3000
	address := fmt.Sprintf(":%d", port)

	route.HealthRoute(app)
	route.WorkflowRoute(app)
	route.LibraryRoute(app)

	go workflow.InitWorker()

	err := app.Listen(address)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error occurred: %+v", err)
		return
	}
}
