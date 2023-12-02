package implementation

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"go.temporal.io/sdk/client"
	"learn-temporal-go/internal/entity/dto"
	"learn-temporal-go/internal/usecase"
	"learn-temporal-go/internal/workflow"
	"learn-temporal-go/internal/workflow/definition"
	"time"
)

type libraryImpl struct {
}

func NewLibraryImpl() usecase.LibraryUseCase {
	return &libraryImpl{}
}

func (l libraryImpl) BorrowBook(request dto.BorrowBookRequestPayload) error {
	log.Info().Msg("Start trigger workflow")

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Error().Msgf("Failed to dial client. Error: %v", err)
		return err
	}
	defer c.Close()

	currentTimestamp := time.Now().UTC().Format("20060102")
	taskId := fmt.Sprintf("borrow_%s_%s", request.UserId, currentTimestamp)

	options := client.StartWorkflowOptions{
		ID:        taskId,
		TaskQueue: workflow.BorrowBookTaskQueue,
	}

	execution, err := c.ExecuteWorkflow(context.Background(), options, definition.BorrowBookWorkflow, request)
	if err != nil {
		log.Error().Msgf("Error occured when executing borrow book sequence. Error: %v", err)
		return err
	}

	err = execution.Get(context.Background(), nil)
	if err != nil {
		log.Error().Msgf("Fetching workflow result failed with error: %v", err)
		return err
	}

	log.Info().Msgf("Workflow with ID: %s and RunID: %s finished!", execution.GetID(), execution.GetRunID())
	return nil
}
