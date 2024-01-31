package ports

import "github.com/kevinkimutai/ticketingapp/payment/application/domain"

type StripePort interface {
	CreateCheckoutSession(payment domain.Payment)
}
