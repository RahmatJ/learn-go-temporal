package activity

import (
	"context"
	"github.com/rs/zerolog/log"
	"learn-temporal-go/internal/entity/dto"
)

func UpdateBorrowRecordStatusActivity(_ context.Context, request dto.UpdateBorrowRecordStatusRequest) error {

	log.Info().Msgf("UPDATED: recordId: %s to status %s", request.RecordId, request.Status)

	return nil
}
