package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"learn-temporal-go/internal/entity/dto"
	"learn-temporal-go/internal/usecase"
)

type library struct {
	uc usecase.LibraryUseCase
}

func NewLibraryHandler(app *fiber.App, uc usecase.LibraryUseCase) {
	handler := &library{
		uc: uc,
	}

	baseGroup := app.Group("/library")
	baseGroup.Post("/borrow", handler.borrowBooks)
	baseGroup.Get("/book/:bookId", handler.getBookById)
}

func (l *library) borrowBooks(ctx *fiber.Ctx) error {
	payload := new(dto.BorrowBookRequestPayload)

	err := ctx.BodyParser(payload)
	if err != nil {
		log.Error().Msgf("Error happened when parsing body %+v", err)
		return err
	}

	log.Info().Msgf("Processing borrow book: %v", *payload)

	err = l.uc.BorrowBook(*payload)

	return ctx.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"data": "Trigger borrow book",
	})
}

func (l *library) getBookById(ctx *fiber.Ctx) error {
	bookId := ctx.Params("bookId")

	log.Info().Msgf("Get Book with BookId: %s", bookId)

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"data": "Get Book Data",
	})
}
