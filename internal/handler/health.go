package handler

import "github.com/gofiber/fiber/v2"

type health struct {
}

func NewHealthHandler(f *fiber.App) {
	handler := &health{}

	f.Get("/ping", handler.healthCheck)
}

func (h *health) healthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(&fiber.Map{
		"data": "PONG",
	})
}
