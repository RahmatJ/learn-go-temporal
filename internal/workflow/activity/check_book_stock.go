package activity

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
)

func CheckBookStockActivity(_ context.Context, bookIds []string) error {

	for _, bookId := range bookIds {
		if bookId == "FAILED_BOOK_ID" {
			log.Error().Msgf("FAILED: Book %s have empty stock", bookId)
			return errors.New("user have borrowed item")
		}

		log.Info().Msgf("SUCCEED: Book %s have available stock", bookId)
	}

	log.Info().Msgf("SUCCEED: All book have available stock")

	return nil
}
