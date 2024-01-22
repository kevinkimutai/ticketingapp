package ports

import "github.com/kevinkimutai/ticketingapp/event/application/domain"

type APIPort interface {
	CreateEvent(event domain.Event) (domain.Event, error)
}
