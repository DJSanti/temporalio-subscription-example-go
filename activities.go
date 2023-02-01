package app

import "context"

// email activity
func SendEmail(_ context.Context, billingInfo BillingInfo, emailPending ComposeEmail) {
	// check for welcome email
	if billingInfo.EmailsSent == 0 {
		// send welcome email
	} else if billingInfo.isCanceled {
		// send cancellation email
	} else if !(billingInfo.isSubscribed) {
		// send parting email
	} else {
		// send subscription email
	}
}