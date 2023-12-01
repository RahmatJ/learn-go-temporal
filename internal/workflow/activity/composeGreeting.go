package activity

import (
	"context"
	"fmt"
)

func ComposeGreeting(_ context.Context, name string) (string, error) {
	greeting := fmt.Sprintf("Hello %s!!! How is it going today?", name)
	return greeting, nil
}
