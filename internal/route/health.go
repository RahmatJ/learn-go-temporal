package route

import (
	"github.com/gofiber/fiber/v2"
	"learn-temporal-go/internal/handler"
)

func HealthRoute(f *fiber.App) {
	handler.NewHealthHandler(f)
}
