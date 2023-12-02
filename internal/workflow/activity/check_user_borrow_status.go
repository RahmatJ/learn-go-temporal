package activity

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
)

func CheckUserBorrowStatusActivity(_ context.Context, userId string) error {
	if userId == "FAILED_USER" {
		log.Error().Msgf("FAILED: User %s have borrowed item", userId)
		return errors.New("user have borrowed item")
	}

	log.Info().Msgf("User %s didn't have borrowed book", userId)

	return nil
}
