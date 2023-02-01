package app

import (
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	// create client and worker
	c, err := client.Dial(client.Options {
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal Client.", err)
	}
	defer c.Close()

	//w := worker.New(c, "SUBSCRIPTION_TASK_QUEUE", worker.Options{})

	// register free trial workflow and subscription workflow

	// register activities
}