package app

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main () {
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client.", err)
	}
	defer c.Close()

	w := worker.New(c, "freeTrialSubscription", worker.Options{})

	w.RegisterWorkflow()
	w.RegisterWorkflow()
	w.RegisterActivity()
	w.RegisterActivity()
	w.RegisterActivity()
	w.RegisterActivity()

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Worker.", err)
	}
}