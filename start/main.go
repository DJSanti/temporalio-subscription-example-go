package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.temporal.io/sdk/client"
)

func main() {
	// create client
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal Client.", err)
	}
	defer c.Close()

	// workflowID format
	workflowID := "SUBSCRIPTION--" + fmt.Sprintf("%d", time.Now().Unix())

	// define workflow options
	options := client.StartWorkflowOptions{
		ID: workflowID,
		TaskQueue: "SUBSCRIPTION_TASK_QUEUE",
	}

	// define wf execution
	we, err := c.ExecuteWorkflow(context.Background(), options, SubscriptionWorkflow, "")
	if err != nil {
		log.Fatalln("Unable to execute Workflow.", err)
	}

	// something to signal here
	// signal
	
	// query workflow
}