package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"learn-temporal-go/internal/entity/dto"
	"learn-temporal-go/internal/usecase"
	wf "learn-temporal-go/internal/workflow"
)

type workflow struct {
	workflowUC usecase.WorkflowUseCase
}

func NewWorkflowHandler(app *fiber.App, useCase usecase.WorkflowUseCase) {
	handler := &workflow{
		workflowUC: useCase,
	}

	baseGroup := app.Group("/workflow")
	baseGroup.Post("/start", handler.startWorkflow)
}

func (w *workflow) startWorkflow(ctx *fiber.Ctx) error {
	payload := new(dto.StartGreetingWorkflowRequestPayload)

	err := ctx.BodyParser(payload)
	if err != nil {
		log.Fatal().Msgf("Bad request %+v", err)
		return err
	}
	log.Info().Msgf("Workflow Initiated ...")

	switch payload.TaskQueue {
	case wf.GreetingTaskQueue:
		greeting, err := w.workflowUC.StartGreeting(payload.Name)
		if err != nil {
			log.Fatal().Msgf("Error when start workflow %+v", err)
			return err
		}

		return ctx.JSON(&fiber.Map{
			"data": greeting,
		})

	case wf.CalculationTaskQueue:
		calculation, err := w.workflowUC.StartCalculate(payload.Number)
		if err != nil {
			log.Fatal().Msgf("Error when start workflow %+v", err)
			return err
		}

		return ctx.JSON(&fiber.Map{
			"data": calculation,
		})

	default:
		log.Fatal().Msgf("TaskQueue type not found %+v", payload.TaskQueue)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Unknown task queue type")
	}
}
