package stripepkg

import (
	"github.com/charmbracelet/log"
	"github.com/kevinkimutai/ticketingapp/payment/application/domain"

	//"github.com/stripe/stripe-go/v72/checkout/session"
	"github.com/stripe/stripe-go/v76"
)

type Adapter struct {
	stripekey string
}

func NewAdapter(stripeSecret string) *Adapter {
	return &Adapter{stripekey: stripeSecret}
}

func (a Adapter) CreateCheckoutSession(payment domain.Payment) {
	stripe.Key = a.stripekey

	var paymentItems []*stripe.CheckoutSessionLineItemParams

	for _, val := range payment.Items {
		items := convertToStripeTypes(val)
		paymentItems = append(paymentItems, items)
	}

	domain := "http://localhost:3000" //spin up a react server
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		LineItems:  paymentItems,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/success.html"),
		CancelURL:  stripe.String(domain + "/cancel.html"),
	}

	log.Info(params)

	// session, err := session.New(params)
	// if err != nil {
	// 	log.Printf("Failed to create Stripe session: %v", err)
	// 	http.Error(w, "Failed to create Stripe session", http.StatusInternalServerError)
	// 	return
	// }

	// fmt.Fprintf(w, "%s", session.ID)
}

func convertToStripeTypes(order domain.OrderItem) *stripe.CheckoutSessionLineItemParams {
	return &stripe.CheckoutSessionLineItemParams{
		PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
			Currency:   stripe.String("KES"),
			Product:    stripe.String(string(order.TicketID)),
			UnitAmount: stripe.Int64(int64(order.Price) * 100),
		},
		Quantity: stripe.Int64(int64(order.Quantity)),
	}

}
