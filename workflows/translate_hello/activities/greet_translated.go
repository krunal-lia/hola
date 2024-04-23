package activities

import (
	"context"
)

func GreetInSelectedLanguage(_ context.Context, greeting string) (string, error) {
	return greeting, nil
}
