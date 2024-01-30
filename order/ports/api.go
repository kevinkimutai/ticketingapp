package ports

import "github.com/kevinkimutai/ticketingapp/order/application/domain"

type APIPort interface {
	CreateNewOrder(domain.Order) (domain.Order, error)
	Verify(token string) (uint64, error)
}
