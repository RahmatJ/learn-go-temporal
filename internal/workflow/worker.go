package workflow

import (
	"github.com/rs/zerolog/log"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"learn-temporal-go/internal/workflow/activity"
	"learn-temporal-go/internal/workflow/definition"
	"sync"
)

func InitWorker() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal().Msgf("unable to create temporal client: %+v", err)
		return
	}
	defer c.Close()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		//Define workflow and activity for each task queue name
		w := worker.New(c, GreetingTaskQueue, worker.Options{})
		w.RegisterWorkflow(definition.GreetingWorkflow)
		w.RegisterActivity(activity.ComposeGreeting)

		err = w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatal().Msgf("unable to start worker: %+v", err)
			return
		}
	}()

	go func() {
		defer wg.Done()

		//Define workflow and activity for each task queue name
		w := worker.New(c, CalculationTaskQueue, worker.Options{})
		w.RegisterWorkflow(definition.CalculationWorkflow)
		w.RegisterActivity(activity.CalculateSquare)

		err = w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatal().Msgf("unable to start worker: %+v", err)
			return
		}
	}()

	go func() {
		defer wg.Done()

		//Define workflow and activity for each task queue name
		w := worker.New(c, BorrowBookTaskQueue, worker.Options{})
		w.RegisterWorkflow(definition.BorrowBookWorkflow)

		w.RegisterActivity(activity.CheckBookStockActivity)
		w.RegisterActivity(activity.CheckUserBorrowStatusActivity)
		w.RegisterActivity(activity.CreateBorrowRecordActivity)
		w.RegisterActivity(activity.UpdateBorrowRecordStatusActivity)
		w.RegisterActivity(activity.UpdateBookStockActivity)
		w.RegisterActivity(activity.UpdateUserBorrowIdActivity)

		err = w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatal().Msgf("unable to start worker: %+v", err)
			return
		}
	}()

	wg.Wait()
}
