package main

import (
	"context"
	"fmt"
	app "hola"
	"hola/workflows/translate_hello"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {

	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "say-hello",
		TaskQueue: app.GreetingTaskQueue,
	}

	var name string
	fmt.Print("What do people call you?\n")
	_, scanErr := fmt.Scan(&name)
	if scanErr != nil {
		log.Fatalln("unable to scan name", err)
	}
	// Start the Workflow
	we, err := c.ExecuteWorkflow(context.Background(), options, translate_hello.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	var language string
	fmt.Print("Which language would you like to be greeted in?\n")
	_, scanErr = fmt.Scan(&language)
	if scanErr != nil {
		log.Fatalln("unable to scan language", err)
	}

	err = c.SignalWorkflow(context.Background(), options.ID, we.GetRunID(), "input-language", language)
	if err != nil {
		log.Fatalln("Error sending the Signal", err)
		return
	}

	// Get the results
	var tranlatedGreeting string
	err = we.Get(context.Background(), &tranlatedGreeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResults(tranlatedGreeting, we.GetID(), we.GetRunID())
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
