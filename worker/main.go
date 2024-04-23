package main

import (
	app "hola"
	"hola/workflows/translate_hello"
	"hola/workflows/translate_hello/activities"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	// This worker hosts both Workflow and Activity functions
	w := worker.New(c, app.GreetingTaskQueue, worker.Options{})
	w.RegisterWorkflow(translate_hello.GreetingWorkflow)
	w.RegisterActivity(activities.GreetEnglish)
	w.RegisterActivity(activities.TranslateThisTo)
	w.RegisterActivity(activities.GreetInSelectedLanguage)

	// Start listening to the Task Queue
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
