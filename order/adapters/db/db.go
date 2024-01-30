package db

import (
	"fmt"

	"github.com/kevinkimutai/ticketingapp/order/application/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID        uint64
	Items         []OrderItem
	TotalAmount   float64
	Currency      string
	PaymentStatus string
	PaymentMethod string
}

type OrderItem struct {
	TicketID uint64
	Quantity uint64
	Price    float64
	Total    float64
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dbString string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	err := db.AutoMigrate(&Order{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a Adapter) CreateOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Create(&order).Error

	if err != nil {
		errMsg := fmt.Sprintf("error creating order :", err)
		return order, status.Errorf(codes.Internal, errMsg)
	}

	return order, nil
}
