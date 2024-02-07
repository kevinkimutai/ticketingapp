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

func (a Application) PaymentRequest(payment domain.Payment) (domain.Payment, error) {
	///Stripe Payment
	stripeID, err := a.stripe.CreateCheckoutSession(payment)
	if err != nil {
		return payment, err
	}

	//Payment struct
	newPayment, err := domain.NewPayment(payment, stripeID)
	if err != nil {
		return payment, err
	}

	//Save To DB
	pay, err := a.db.CreatePayment(newPayment)
	if err != nil {
		return pay, err
	}

	return pay, nil

	//Send Request Saving To Order

	//Ticket Generation
}
