package api

import (
	"github.com/kevinkimutai/ticketingapp/event/application/domain"
	"github.com/kevinkimutai/ticketingapp/event/ports"
)

type Application struct {
	db   ports.DBPort
	auth ports.AuthPort
}

func NewApplication(db ports.DBPort, auth ports.AuthPort) *Application {
	return &Application{db: db, auth: auth}
}

func (a *Application) CreateEvent(event domain.Event) (domain.Event, error) {

	result, err := a.db.Create(event)
	return result, err

}

func (a *Application) Verify(token string) error {
	err := a.auth.Verify(token)

	return err

}
