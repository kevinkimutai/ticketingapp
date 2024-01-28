package api

import (
	"github.com/kevinkimutai/ticketingapp/ticket/application/domain"
	"github.com/kevinkimutai/ticketingapp/ticket/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a Application) CreateTicket(ticket domain.Ticket) (domain.Ticket, error) {
	ticket, err := a.db.CreateEventTicket(ticket)

	return ticket, err
}
