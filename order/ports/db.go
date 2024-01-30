package ports

import "github.com/kevinkimutai/ticketingapp/order/application/domain"

type DBPort interface {
	CreateOrder(domain.Order) (domain.Order, error)
}
