package usecase

import "learn-temporal-go/internal/entity/dto"

type LibraryUseCase interface {
	BorrowBook(request dto.BorrowBookRequestPayload) error
}
