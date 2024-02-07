package domain

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Payment struct {
	ID          uint64      `json:"id"`
	OrderID     uint64      `json:"order_id"`
	UserID      uint64      `json:"user_id"`
	Items       []OrderItem `json:"items"`
	TotalAmount float64     `json:"total_amount"`
	Currency    string      `json:"currency"`
	StripeID    string      `json:"stripe_id"`
}

type OrderItem struct {
	TicketID uint64  `json:"ticket_id"`
	Quantity uint64  `json:"quantity"`
	Price    float64 `json:"price"`
	Total    float64 `json:"total"`
}

func NewPayment(payment Payment, stripeID string) (Payment, error) {
	if stripeID == "" {
		return payment, status.Errorf(codes.InvalidArgument, "Missing StripeID")
	}
	newPayment := Payment{
		OrderID:     payment.OrderID,
		UserID:      payment.UserID,
		StripeID:    stripeID,
		TotalAmount: payment.TotalAmount,
	}
	return newPayment, nil
}
