package definition

import (
	"github.com/pborman/uuid"
	"github.com/rs/zerolog/log"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/multierr"
	"learn-temporal-go/internal/entity/dto"
	"learn-temporal-go/internal/workflow/activity"
	"time"
)

func BorrowBookWorkflow(ctx workflow.Context, request dto.BorrowBookRequestPayload) error {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	context := workflow.WithActivityOptions(ctx, options)

	//	Check user borrow status
	err := workflow.ExecuteActivity(context, activity.CheckUserBorrowStatusActivity, request.UserId).Get(context, nil)
	if err != nil {
		log.Error().Msgf("Error occurred when check borrow status. Error: %v", err)
		return err
	}

	//  Check book stock
	err = workflow.ExecuteActivity(context, activity.CheckBookStockActivity, request.BookIds).Get(context, nil)
	if err != nil {
		log.Error().Msgf("Error occurred when check book stock. Error: %v", err)
		return err
	}

	//  Update book stock
	updateBookStockRequest := dto.UpdateManyBookStockRequest{
		BookIds:    request.BookIds,
		IsIncrease: true,
		Count:      1,
	}
	err = workflow.ExecuteActivity(
		context, activity.UpdateBookStockActivity, updateBookStockRequest,
	).Get(context, nil)
	if err != nil {
		log.Error().Msgf("Error occurred when update book stock. Error: %v", err)
		return err
	}

	defer func() {
		if err != nil {
			updateBookStockRequest = dto.UpdateManyBookStockRequest{
				BookIds:    request.BookIds,
				IsIncrease: false,
				Count:      1,
			}
			errCompensation := workflow.ExecuteActivity(
				context, activity.UpdateBookStockActivity, updateBookStockRequest,
			).Get(context, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()

	//  Create record in borrow_transaction
	createBorrowRecord := dto.CreateBorrowRecordRequest{
		UserId: request.UserId,
		Books:  request.BookIds,
	}

	//generate uuid that idempotent, even after retry
	generateUUID := workflow.SideEffect(ctx, func(_ workflow.Context) interface{} {
		return uuid.NewUUID().String()
	})

	var id string
	err = generateUUID.Get(&id)
	if err != nil {
		log.Error().Msgf("Error when generating UUID. Error: %v", err)
		return err
	}
	//END generate uuid

	var recordId string
	err = workflow.ExecuteActivity(
		context, activity.CreateBorrowRecordActivity, id, createBorrowRecord,
	).Get(context, &recordId)
	if err != nil {
		log.Error().Msgf("Error occurred when create borrow record. Error: %v", err)
		return err
	}

	defer func() {
		if err != nil {
			updateRecordStatus := dto.UpdateBorrowRecordStatusRequest{
				RecordId: recordId,
				Status:   "DROPPED",
			}
			errCompensation := workflow.ExecuteActivity(
				context, activity.UpdateBorrowRecordStatusActivity, updateRecordStatus,
			).Get(context, nil)
			err = multierr.Append(err, errCompensation)
		}
	}()

	//  Update user.borrowingId
	updateUser := dto.UpdateUserBorrowId{
		UserId:   request.UserId,
		BorrowId: recordId,
	}
	err = workflow.ExecuteActivity(
		context, activity.UpdateUserBorrowIdActivity, updateUser,
	).Get(context, nil)
	if err != nil {
		log.Error().Msgf("Error occurred when update user record. Error: %v", err)
		return err
	}

	//  Return succeed
	log.Info().Msgf("Succeed Execute borrow book sequence")
	return nil
}
