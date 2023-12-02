package activity

import (
	"context"
	"github.com/rs/zerolog/log"
	"learn-temporal-go/internal/entity/dto"
)

func CreateBorrowRecordActivity(_ context.Context, id string, request dto.CreateBorrowRecordRequest) (string, error) {

	log.Info().Msgf("Successfully create borrow record for user %s with id: %s", request.UserId, id)

	return id, nil
}
