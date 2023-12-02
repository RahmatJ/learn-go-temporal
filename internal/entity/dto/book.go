package dto

type (
	UpdateBookStockRequest struct {
		BookId     string
		IsIncrease bool
		Count      int64
	}

	UpdateManyBookStockRequest struct {
		BookIds    []string
		IsIncrease bool
		Count      int64
	}
)
