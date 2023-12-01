package usecase

type WorkflowUseCase interface {
	StartGreeting(name string) (string, error)
	StartCalculate(data int64) (int64, error)
}
