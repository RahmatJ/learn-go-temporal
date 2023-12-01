package route

import (
	"github.com/gofiber/fiber/v2"
	"learn-temporal-go/internal/handler"
	"learn-temporal-go/internal/usecase/implementation"
)

func WorkflowRoute(app *fiber.App) {
	workflowUC := implementation.NewWorkflowImpl()
	handler.NewWorkflowHandler(app, workflowUC)
}
