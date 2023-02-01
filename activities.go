package app

import "fmt"

// email activity
func SendEmail(billingInfo BillingInfo, pendingEmail ComposeEmail) {
	// check for welcome email
	if billingInfo.EmailsSent == 0 {
		// send welcome email
		fmt.Println("Sending welcome email to %s with the following message:\n%s", billingInfo.Email, pendingEmail.Message)
		billingInfo.EmailsSent++
		billingInfo.MaxBillingPeriods = 12
		billingInfo.isSubscribed = true
	}  else if !(billingInfo.isSubscribed) || (billingInfo.MaxBillingPeriods == billingInfo.EmailsSent) {
		// send parting email
		fmt.Println("Sending welcome email to %s with the following message:\n%s", billingInfo.Email, pendingEmail.Message)
		billingInfo.isSubscribed = false
	} else {
		// send subscription email
		fmt.Println("Sending subscription update email to %s with the following message:\n%s", billingInfo.Email, pendingEmail.Message)
		billingInfo.EmailsSent++
	}
}