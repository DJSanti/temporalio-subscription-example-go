package app

//TODO: make this more email-centric
type Activities struct {
	Name string
	Greeting string
}

// WelcomeEmail
func (a *Activities) SendWelcomeEmail() error {
	return err
}

// TrialCancellationEmail
func (a *Activities) TrialCancellationEmail() error {
	return err
}

// SubscriptionCancellationEmail
func (a *Activities) SubscriptionCancellationEmail() error {
	return err
}

// SubscriptionEndedEmail
func (a *Activities) SubscriptionEndedEmail() error {
	return err
}