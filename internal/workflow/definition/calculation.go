package definition

import (
	"github.com/rs/zerolog/log"
	"go.temporal.io/sdk/workflow"
	"learn-temporal-go/internal/workflow/activity"
	"time"
)

func CalculationWorkflow(ctx workflow.Context, data int64) (int64, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result int64
	err := workflow.ExecuteActivity(ctx, activity.CalculateSquare, data).Get(ctx, &result)

	if err != nil {
		log.Fatal().Msgf("Error happened: %+v", err)
	}

	return result, err
}
