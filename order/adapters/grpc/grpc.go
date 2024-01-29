package grpc

import (
	"context"

	"github.com/kevinkimutai/ticketingapp/order/application/domain"
	orderproto "github.com/kevinkimutai/ticketingapp/order/proto/golang/order"
)

func (a Adapter) CreateOrder(ctx context.Context, req *orderproto.CreateOrderRequest) (*orderproto.CreateOrderResponse, error) {
	var items []domain.OrderItem

	for _, order := range req.Items {
		domainOrder := ConvertProtoRequestToDomain(order)
		items = append(items, domainOrder)
	}

	userId := ctx.Value("userid").(uint64)

	request := domain.Order{UserID: userId, Items: items}

	newOrder, err := domain.NewOrder(request)
	if err != nil {
		return nil, err
	}

	a.api.CreateNewOrder()

}

func ConvertProtoRequestToDomain(order *orderproto.OrderItems) domain.OrderItem {
	domainOrder := domain.OrderItem{
		TicketID: order.TicketId,
		Quantity: order.Quantity,
		Price:    float64(order.Price),
		Total:    float64(order.Total),
	}
	return domainOrder
}
