package main

import (
	"github.com/rs/zerolog/log"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"learn-temporal-go/internal/workflow"
	"learn-temporal-go/internal/workflow/activity"
	"learn-temporal-go/internal/workflow/definition"
	"sync"
)

func main() {
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
		w := worker.New(c, workflow.GreetingTaskQueue, worker.Options{})
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
		w := worker.New(c, workflow.CalculationTaskQueue, worker.Options{})
		w.RegisterWorkflow(definition.CalculationWorkflow)
		w.RegisterActivity(activity.CalculateSquare)

		err = w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatal().Msgf("unable to start worker: %+v", err)
			return
		}
	}()

	wg.Wait()
}
