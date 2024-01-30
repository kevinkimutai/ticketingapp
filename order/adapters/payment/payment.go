package payment

import (
	"context"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/kevinkimutai/ticketingapp/order/application/domain"
	paymentproto "github.com/kevinkimutai/ticketingapp/order/proto/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment paymentproto.PaymentProtoClient
}

func NewAdapter(organiserServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
	)))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(organiserServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := paymentproto.NewPaymentProtoClient(conn)
	return &Adapter{payment: client}, nil
}

func (a Adapter) CreatePayment(order domain.Order) (uint64, error) {

	//Convert Items to ]*paymentproto.OrderItems
	var items []*paymentproto.OrderItems

	for _, o := range order.Items {
		paymentOrder := ConvertOrderItemsToPaymentProtoItems(o)
		items = append(items, paymentOrder)
	}

	response, err := a.payment.CreatePayment(context.Background(),
		&paymentproto.CreatePaymentRequest{
			OrderId:     order.ID,
			UserId:      order.UserID,
			Items:       items,
			TotalAmount: float32(order.TotalAmount),
			Currency:    order.Currency,
		})

	return response.PaymentId, err
}

func ConvertOrderItemsToPaymentProtoItems(item domain.OrderItem) *paymentproto.OrderItems {
	protoVal := &paymentproto.OrderItems{
		TicketId: item.TicketID,
		Quantity: item.Quantity,
		Price:    float32(item.Price),
		Total:    float32(item.Total),
	}
	return protoVal

}
