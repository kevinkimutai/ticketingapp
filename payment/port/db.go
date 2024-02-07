package ports

import "github.com/kevinkimutai/ticketingapp/payment/application/domain"

type DBPort interface {
	CreatePayment(payment domain.Payment) (domain.Payment, error)
}
