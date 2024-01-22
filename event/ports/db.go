package ports

import "github.com/kevinkimutai/ticketingapp/event/application/domain"

type DBPort interface {
	Create(domain.Event) (domain.Event, error)
}
