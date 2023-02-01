package app

import (
	"github.com/mitchellh/mapstructure"
	"go.temporal.io/sdk/workflow"
)

// Workflow definition
func SubscriptionWorkflow(ctx workflow.Context) {
	// set up logger
	logger := workflow.GetLogger(ctx)

	// declare structs here
	var PendingEmail ComposeEmail
	var billingInfo BillingInfo

	// set up query handler
	err := workflow.SetQueryHandler(ctx, "getBillingInfo", func(input []byte) (billingInfo, error) {	
		if err != nil {
			logger.Info("SetQueryHandler failed.", "Error", err)
		}
	})

	// set up signal channel(s)
	welcomeEmailChannel := workflow.GetSignalChannel(ctx, SignalChannels.WELCOME_EMAIL)
	cancelFreeTrialChannel := workflow.GetSignalChannel(ctx, SignalChannels.CANCEL_FREE_EMAIL)
	cancelSubscriptionChannel := workflow.GetSignalChannel(ctx, SignalChannels.CANCEL_SUBSCRIPTION_EMAIL)
	subscriptionEndChannel := workflow.GetSignalChannel(ctx, SignalChannels.SUBSCRIPTION_ENDED_EMAIL)

	isSubscribed := true
	
	// workflow logic
	for (isSubscribed) {
		selector := workflow.NewSelector(ctx)
		// signal handler for welcome email
		selector.AddReceive(welcomeEmailChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)

			PendingEmail.Message = "Welcome to your new subscription! This is the very first email and the start of the trial period!"

			err := mapstructure.Decode(signal, &PendingEmail)
			if err != nil {
				logger.Error("Invalid Signal type: %v", err)
				return
			}
			billingInfo.Email = "example@google.co"
			billingInfo.EmailsSent = 0

			SendEmail(billingInfo, PendingEmail)
			isSubscribed = billingInfo.isSubscribed
		})
		// signal handler for cancellation
		selector.AddReceive(cancelFreeTrialChannel, func(c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)

			PendingEmail.Message = "We're so sorry to see you go—and during your trial period, no less! You will no longer receive any emails from us. Goodbye!"

			err := mapstructure.Decode(signal, &PendingEmail)
			if err != nil {
				logger.Error("Invalid Signal type: %v", err)
				return
			}
			SendEmail(billingInfo, PendingEmail)
			isSubscribed = billingInfo.isSubscribed

		})
			// signal handler for expired subscription
		selector.AddReceive(cancelSubscriptionChannel, func (c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)
			PendingEmail.Message = "We're so sorry to see you go—and during your trial period, no less! You will no longer receive any emails from us. Goodbye!"

			err := mapstructure.Decode(signal, &PendingEmail)
			if err != nil {
				logger.Error("Invalid Signal type: %v", err)
				return
			}
			SendEmail(billingInfo, PendingEmail)
			isSubscribed = billingInfo.isSubscribed
		})
		selector.AddReceive(subscriptionEndChannel, func (c workflow.ReceiveChannel, _ bool) {
			var signal interface{}
			c.Receive(ctx, &signal)
			PendingEmail.Message = "It appears that your subscription has come to an end. Farewell!"

			err := mapstructure.Decode(signal, &PendingEmail)
			if err != nil {
				logger.Error("Invalid Signal type: %v", err)
				return
			}
			SendEmail(billingInfo, PendingEmail)
			isSubscribed = billingInfo.isSubscribed
		})
	}
}



	
