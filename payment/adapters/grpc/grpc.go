package grpc

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/payment/application/domain"
	paymentproto "github.com/kevinkimutai/ticketingapp/payment/proto/golang/payment"
)

func (a Adapter) CreatePayment(ctx context.Context, req *paymentproto.CreatePaymentRequest) (*paymentproto.CreatePaymentResponse, error) {
	//Convert To Domain.Payment
	var items []domain.OrderItem

	for _, item := range req.Items {
		i := convertProtoRequestToDomain(item)
		items = append(items, i)

	}

	request := domain.Payment{
		OrderID:     req.OrderId,
		UserID:      req.UserId,
		Items:       items,
		TotalAmount: float64(req.TotalAmount),
		Currency:    req.Currency,
	}

	payment, err := a.api.PaymentRequest(request)
	if err != nil {
		return nil, err
	}

	return &paymentproto.CreatePaymentResponse{
		PaymentId: payment.ID,
	}, nil
}

func convertProtoRequestToDomain(items *paymentproto.OrderItems) domain.OrderItem {
	return domain.OrderItem{
		TicketID: items.TicketId,
		Quantity: items.Quantity,
		Price:    float64(items.Price),
		Total:    float64(items.Total),
	}
}
