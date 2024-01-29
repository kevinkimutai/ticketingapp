package ports

import (
	"github.com/kevinkimutai/ticketingapp/event/application/domain"
)

type APIPort interface {
	CreateEvent(uint64, domain.Event) (domain.Event, error)
	GetAllEvents(params domain.Params) ([]domain.Event, error)
}
