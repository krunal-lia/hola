package translate_hello

import (
	"go.temporal.io/sdk/temporal"
	"hola/workflows/translate_hello/activities"
	"time"

	"go.temporal.io/sdk/workflow"
)

func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	retryPolicy := &temporal.RetryPolicy{
		MaximumAttempts:        1,
		NonRetryableErrorTypes: []string{"LanguageNotFound"},
	}

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
		RetryPolicy:         retryPolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	greetErr := workflow.ExecuteActivity(ctx, activities.GreetEnglish, name).Get(ctx, &result)
	if greetErr != nil {
		return "", greetErr
	}

	var language string
	signalChan := workflow.GetSignalChannel(ctx, "input-language")
	signalChan.Receive(ctx, &language)

	var greetingInSelectedLanguage string
	translateErr := workflow.ExecuteActivity(ctx, activities.TranslateThisTo, language).Get(ctx, &greetingInSelectedLanguage)

	var finalGreeting string
	if translateErr != nil {
		finalGreeting = "Sorry I don't know this language"
	} else {
		finalGreeting = greetingInSelectedLanguage + " " + name + "!"
	}

	var finalResult string
	err := workflow.ExecuteActivity(ctx, activities.GreetInSelectedLanguage, finalGreeting).Get(ctx, &finalResult)
	if err != nil {
		return "", err
	}

	return finalGreeting, nil
}
