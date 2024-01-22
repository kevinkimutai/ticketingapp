package api

import (
	"github.com/kevinkimutai/ticketingapp/event/application/domain"
	"github.com/kevinkimutai/ticketingapp/event/ports"
)

type Application struct {
	db  ports.DBPort
	api ports.APIPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) CreateEvent(event domain.Event) (domain.Event, error) {
	
	result, err := a.db.Create(event)
	return result, err

}
