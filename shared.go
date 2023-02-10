package app

type BillingInfo struct {
	Email string
	MaxBillingPeriods int
	EmailsSent int
	isSubscribed bool
}

type ComposeEmail struct {
	Route string
	Message string
}