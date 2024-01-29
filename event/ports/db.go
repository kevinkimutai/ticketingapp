package ports

import (
	"github.com/kevinkimutai/ticketingapp/event/application/domain"
	"gorm.io/gorm"
)

type DBPort interface {
	Create(domain.Event) (domain.Event, error)
	BeginTx() *gorm.DB
	GetEvents(domain.Params) ([]domain.Event, error)
}
