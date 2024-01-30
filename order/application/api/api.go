package api

import (
	"github.com/kevinkimutai/ticketingapp/order/application/domain"
	"github.com/kevinkimutai/ticketingapp/order/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
	auth    ports.AuthPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, auth ports.AuthPort) *Application {
	return &Application{db: db, payment: payment, auth: auth}
}

func (a Application) Verify(token string) (uint64, error) {
	userId, err := a.auth.Verify(token)

	return userId, err

}

func (a Application) CreateNewOrder(order domain.Order) (domain.Order, error) {

	//Create Order
	newOrder, err := a.db.CreateOrder(order)
	if err != nil {
		return newOrder, err
	}

	//Create Payment Intent
	a.payment.CreatePayment()

	//If success
	//Change status To Paid

	//Create Ticket Service
}
