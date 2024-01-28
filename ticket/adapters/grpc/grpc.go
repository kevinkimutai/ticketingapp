package grpc

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/ticket/application/domain"
	ticketproto "github.com/kevinkimutai/ticketingapp/ticket/proto/golang"
)

func (a Adapter) CreateTicket(ctx context.Context, req *ticketproto.CreateTicketRequest) (*ticketproto.CreateTicketResponse, error) {

	request := domain.Ticket{TicketType: req.TicketType, NumberOfTickets: req.NumberOfTickets, Price: float64(req.Price), EventID: req.EventId}

	ticket, err := domain.NewTicket(request)
	if err != nil {
		return nil, err
	}

	newTicket, err := a.api.CreateTicket(ticket)
	if err != nil {
		return nil, err
	}

	return &ticketproto.CreateTicketResponse{ID: newTicket.ID}, nil
}
