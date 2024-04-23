package activities

import (
	"context"
	"fmt"
)

func GreetEnglish(_ context.Context, name string) (string, error) {
	greeting := fmt.Sprintf("Hello %s!", name)
	return greeting, nil
}
