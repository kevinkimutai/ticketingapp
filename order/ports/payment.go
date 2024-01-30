package ports

type Payment struct{
	
}
type PaymentPort interface {
	CreatePayment() (uint64, error)
}
