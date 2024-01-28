package db

import (
	"fmt"

	"github.com/kevinkimutai/ticketingapp/ticket/application/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	TicketTypeID    int64
	NumberOfTickets uint64
	Price           float64
	EventID         int64
	TicketType      string
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dbString string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	err := db.AutoMigrate(&Ticket{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a *Adapter) CreateEventTicket(ticket domain.Ticket) (domain.Ticket, error) {
	err := a.db.Create(&ticket).Error

	if err != nil {
		errMsg := fmt.Sprintf("Internal Error:%d", err)
		return ticket, status.Errorf(codes.Internal, errMsg)
	}

	return ticket, nil
}
