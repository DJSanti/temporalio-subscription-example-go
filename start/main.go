package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.temporal.io/sdk/client"

	"temporalio-subscription-example-go/app"
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
	we, err := c.ExecuteWorkflow(context.Background(), options, app.SubscriptionWorkflow, "")
	if err != nil {
		log.Fatalln("Unable to execute Workflow.", err)
	}

	// signal welcome email
	update := app.ComposeEmail{Route: app.SignalChannels.WELCOME_EMAIL, Message: "testing..."}
	err = c.SignalWorkflow(context.Background(), we.GetID(), "", "SEND_WELCOME_EMAIL", update)
	// query workflow
	response, err := c.QueryWorkflow(context.Background(), we.GetID(), "", "getBillingInfo")
	if err != nil {
		log.Fatalln("Unable to query Workflow.", err)
	}
	var result interface{}
	if err := response.Get(&result); err != nil {
		log.Fatalln("Unable to decode Query result.", err)
	}
	log.Println("Received Query result.", "Result:", result)

	// signals until subscription runs out
}