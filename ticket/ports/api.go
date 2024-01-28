package ports

import "github.com/kevinkimutai/ticketingapp/ticket/application/domain"

type APIPort interface {
	CreateTicket(ticket domain.Ticket) (domain.Ticket, error)
}
