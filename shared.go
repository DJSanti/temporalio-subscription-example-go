package app


var SignalChannels = struct {
	WELCOME_EMAIL string
	CANCEL_FREE_EMAIL string
	CANCEL_SUBSCRIPTION_EMAIL string
	SUBSCRIPTION_ENDED_EMAIL string
}{
	WELCOME_EMAIL:  "SEND_WELCOME_EMAIL",
	CANCEL_FREE_EMAIL: "SEND_CANCEL_FREE_EMAIL",
	CANCEL_SUBSCRIPTION_EMAIL: "SEND_CANCEL_SUBSCRIPTION_EMAIL",
	SUBSCRIPTION_ENDED_EMAIL: "SEND_SUBSCRIPTION_ENDED_EMAIL",
}

type BillingInfo struct {
	Email string
	MaxBillingPeriods int
	EmailsSent int
	isSubscribed bool
}

type ComposeEmail struct {
	Message string
}