package domain

type Payment struct {
	ID          uint64      `json:"id"`
	OrderID     uint64      `json:"order_id"`
	UserID      uint64      `json:"user_id"`
	Items       []OrderItem `json:"items"`
	TotalAmount float64     `json:"total_amount"`
	Currency    string      `json:"currency"`
}

type OrderItem struct {
	TicketID uint64  `json:"ticket_id"`
	Quantity uint64  `json:"quantity"`
	Price    float64 `json:"price"`
	Total    float64 `json:"total"`
}
