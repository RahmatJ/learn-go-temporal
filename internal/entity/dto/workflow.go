package dto

type (
	StartGreetingWorkflowRequestPayload struct {
		Name      string `json:"name"`
		TaskQueue string `json:"taskQueue"`
		Number    int64  `json:"number"`
	}
)
