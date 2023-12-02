package route

import (
	"github.com/gofiber/fiber/v2"
	"learn-temporal-go/internal/handler"
	"learn-temporal-go/internal/usecase/implementation"
)

func LibraryRoute(app *fiber.App) {
	uc := implementation.NewLibraryImpl()
	handler.NewLibraryHandler(app, uc)
}
