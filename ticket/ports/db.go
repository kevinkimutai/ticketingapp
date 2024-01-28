package ports

import "github.com/kevinkimutai/ticketingapp/ticket/application/domain"

type DBPort interface {
	CreateEventTicket(ticket domain.Ticket) (domain.Ticket, error)
}
