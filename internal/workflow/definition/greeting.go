package definition

import (
	"github.com/rs/zerolog/log"
	"go.temporal.io/sdk/workflow"
	"learn-temporal-go/internal/workflow/activity"
	"time"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, activity.ComposeGreeting, name).Get(ctx, &result)

	if err != nil {
		log.Fatal().Msgf("Error happened: %+v", err)
	}

	return result, err
}
