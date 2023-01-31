package app

import (
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal Client.", err)
	} 
	defer c.Close()

	workflowID := ""

	options := client.StartWorkflowOptions{
		ID: "",
		TaskQueue:"",
	}

	// state of customer??

	// Signal Handler

	// Query

	// interface

	//print result
}