package app

import (
	"go.temporal.io/sdk/workflow"
)

type ComposeEmail struct {
	Email string
	Message string
	count int
}

type BillingInfo struct {
	Email string
	MaxBillingPeriods int
	EmailsSent int
	isSubscribed bool
	isCanceled bool

}

// FreeTrialWorkflow definition
func SubscriptionWorkflow(ctx workflow.Context) {
	// set up logger
	logger := workflow.GetLogger(ctx)

	// declare struct here
	var PendingEmail ComposeEmail

	// set up query handler
	err := workflow.SetQueryHandler(ctx, "getBillingInfo", "") 
	if err != nil {
		logger.Info("SetQueryHandler failed.", "Error", err)
	}

	// set up signal channel(s)

	// workflow logic here

	// signal handler for welcome email

	// signal handler for cancellation

	// signal handler for expired subscription

	// check billing period for expiration

}

// set free trial state

// set subscription state



	
