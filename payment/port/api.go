package ports

import "github.com/kevinkimutai/ticketingapp/payment/application/domain"

type APIPort interface {
	PaymentRequest(payment domain.Payment) (domain.Payment, error)
	Verify(token string) (uint64, error)
}
