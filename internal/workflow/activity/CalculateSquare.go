package activity

import "context"

func CalculateSquare(_ context.Context, data int64) (int64, error) {
	result := data * data
	return result, nil
}
