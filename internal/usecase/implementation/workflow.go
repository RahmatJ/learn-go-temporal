package implementation

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.temporal.io/sdk/client"
	"learn-temporal-go/internal/usecase"
	"learn-temporal-go/internal/workflow"
	"learn-temporal-go/internal/workflow/definition"
)

type workflowImpl struct {
}

func NewWorkflowImpl() usecase.WorkflowUseCase {
	return &workflowImpl{}
}

func (w workflowImpl) StartGreeting(name string) (string, error) {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal().Msgf("unable to call workflow: %+v", err)
		return "", err
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: workflow.GreetingTaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, definition.GreetingWorkflow, name)
	if err != nil {
		log.Fatal().Msgf("unable to execute workflow: %+v", err)
		return "", err
	}

	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatal().Msgf("unable to get workflow result: %+v", err)
		return "", err
	}

	log.Info().Msgf("Workflow with ID: %s and RunID: %s finished!", we.GetID(), we.GetRunID())
	return greeting, nil
}

func (w workflowImpl) StartCalculate(data int64) (int64, error) {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal().Msgf("unable to call workflow: %+v", err)
		return -1, err
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "calculation-workflow",
		TaskQueue: workflow.CalculationTaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, definition.CalculationWorkflow, data)
	if err != nil {
		log.Fatal().Msgf("unable to execute workflow: %+v", err)
		return -1, err
	}

	var calculation int64
	err = we.Get(context.Background(), &calculation)
	if err != nil {
		log.Fatal().Msgf("unable to get workflow result: %+v", err)
		return -1, err
	}

	log.Info().Msgf("Workflow with ID: %s and RunID: %s finished!", we.GetID(), we.GetRunID())
	return calculation, nil
}
