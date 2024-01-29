package db

import (
	"fmt"

	"github.com/kevinkimutai/ticketingapp/event/application/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name         string
	Venue        string
	Town         string
	Longitude    float64
	Latitude     float64
	PosterImgURL string `json:"poster_img"`
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dbString string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	err := db.AutoMigrate(&Event{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a *Adapter) Create(event domain.Event) (domain.Event, error) {

	err := a.db.Create(&event).Error

	if err != nil {
		return event, err
	}

	return event, nil
}

func (a *Adapter) BeginTx() *gorm.DB {
	return a.db.Begin()
}

func (a *Adapter) GetEvents(params domain.Params) ([]domain.Event, error) {
	events := []domain.Event{}

	offset := (params.PageNumber - 1) * params.PageSize
	err := a.db.Offset(int(offset)).Limit(int(params.PageSize)).Find(&events).Error

	if err != nil {
		errMsg := fmt.Sprintf("something went wrong %d", err)
		return events, status.Errorf(codes.Internal, errMsg)
	}

	return events, nil
}
