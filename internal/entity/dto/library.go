package dto

type (
	BorrowBookRequestPayload struct {
		UserId  string   `json:"userId"`
		BookIds []string `json:"bookIds"`
	}

	CreateBorrowRecordRequest struct {
		UserId string
		Books  []string
	}

	UpdateBorrowRecordStatusRequest struct {
		RecordId string
		Status   string
	}
)
