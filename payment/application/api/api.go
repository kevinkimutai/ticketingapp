package api

import (
	"github.com/kevinkimutai/ticketingapp/payment/application/domain"
	ports "github.com/kevinkimutai/ticketingapp/payment/port"
)

type Application struct {
	db     ports.DBPort
	stripe ports.StripePort
	auth   ports.AuthPort
}

func NewApplication(db ports.DBPort, stripe ports.StripePort, auth ports.AuthPort) *Application {
	return &Application{db: db, stripe: stripe, auth: auth}
}

func (a Application) Verify(token string) (uint64, error) {
	userId, err := a.auth.Verify(token)

	return userId, err

}

func (a Application) PaymentRequest(payment domain.Payment) {
	///Stripe Payment
	a.stripe.CreateCheckoutSession(payment)

	//Save To DB

	//Send Request Saving To Order

	//Ticket Generation
}
