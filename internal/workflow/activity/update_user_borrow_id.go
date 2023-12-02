package activity

import (
	"context"
	"github.com/rs/zerolog/log"
	"learn-temporal-go/internal/entity/dto"
)

func UpdateUserBorrowIdActivity(_ context.Context, request dto.UpdateUserBorrowId) error {

	log.Info().Msgf("UPDATING: user %s with borrowId %s", request.UserId, request.BorrowId)

	return nil
}
