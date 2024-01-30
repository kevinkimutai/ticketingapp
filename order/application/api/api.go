package api

import (
	"fmt"

	"github.com/kevinkimutai/ticketingapp/order/application/domain"
	"github.com/kevinkimutai/ticketingapp/order/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	_, err = a.payment.CreatePayment(newOrder)
	if err != nil {
		errMsg := fmt.Sprintf("payment request error :%d", err)
		return newOrder, status.Errorf(codes.Internal, errMsg)
	}

	return newOrder, nil

}
