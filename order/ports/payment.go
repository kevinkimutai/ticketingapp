package ports

import "github.com/kevinkimutai/ticketingapp/order/application/domain"

type PaymentPort interface {
	CreatePayment(domain.Order) (uint64, error)
}
