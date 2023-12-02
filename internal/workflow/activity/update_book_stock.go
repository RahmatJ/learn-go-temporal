package activity

import (
	"context"
	"github.com/rs/zerolog/log"
	"learn-temporal-go/internal/entity/dto"
)

func UpdateBookStockActivity(_ context.Context, request dto.UpdateManyBookStockRequest) error {
	for _, bookId := range request.BookIds {
		operationString := "SUBSTRACT"
		if request.IsIncrease {
			operationString = "ADD"
		}
		log.Info().Msgf("UPDATING: bookId: %s, operation: %s, count: %d",
			bookId, operationString, request.Count,
		)
	}

	log.Info().Msg("Success UPDATING records for books")
	return nil
}
