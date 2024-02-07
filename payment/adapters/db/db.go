package db

import (
	"fmt"

	"github.com/kevinkimutai/ticketingapp/payment/application/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	OrderID     uint64
	UserID      uint64
	TotalAmount float64
	StripeID    string
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dbString string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	err := db.AutoMigrate(&Payment{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a Adapter) CreatePayment(payment domain.Payment) (domain.Payment, error) {

	err := a.db.Create(&payment).Error

	if err != nil {
		errMsg := fmt.Sprintf("error creating payment : %d", err)
		return payment, status.Errorf(codes.Internal, errMsg)
	}

	return payment, nil
}
