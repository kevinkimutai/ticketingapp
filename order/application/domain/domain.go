package domain

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Order struct {
	ID            uint64      `json:"id"`
	UserID        uint64      `json:"user_id"`
	Items         []OrderItem `json:"items"`
	TotalAmount   float64     `json:"total_amount"`
	Currency      string      `json:"currency"`
	PaymentStatus string      `json:"paymet_status"`
	PaymentMethod string      `json:"payment_method"`
}

type OrderItem struct {
	TicketID uint64  `json:"ticket_id"`
	Quantity uint64  `json:"quantity"`
	Price    float64 `json:"price"`
	Total    float64 `json:"total"`
}

func NewOrder(o Order) (Order, error) {
	if o.UserID == 0 || o.Items == nil || o.TotalAmount == 0 {
		return o, status.Errorf(codes.InvalidArgument, "missing order inputs")
	}
	o.Currency = "KES"
	o.PaymentMethod = "Credit Card"
	o.PaymentStatus = "Pending"

	return o, nil

}

//TODO:USE REDIS TO CACHE TICKETS
