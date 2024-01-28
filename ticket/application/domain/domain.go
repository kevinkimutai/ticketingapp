package domain

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Ticket struct {
	ID uint64 `json:"id"`

	NumberOfTickets uint64  `json:"number_of_tickets"`
	Price           float64 `json:"price"`
	EventID         int64   `json:"event_id"`
	TicketType      string  `json:"ticket_type"`
}

func NewTicket(ticket Ticket) (Ticket, error) {
	if ticket.EventID == 0 || ticket.TicketType == "" || ticket.NumberOfTickets == 0 || ticket.Price == 0 {
		return ticket, status.Errorf(codes.InvalidArgument, "missing input values in ticket creation")
	}
	return ticket, nil
}
