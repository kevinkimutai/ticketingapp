package api

import (
	"errors"
	"fmt"

	"github.com/kevinkimutai/ticketingapp/event/application/domain"
	"github.com/kevinkimutai/ticketingapp/event/ports"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db        ports.DBPort
	organiser ports.OrganiserPort
}

func NewApplication(db ports.DBPort, organiser ports.OrganiserPort) *Application {
	return &Application{db: db, organiser: organiser}
}

func (a *Application) CreateEvent(userId uint64, event domain.Event) (domain.Event, error) {

	//Begin TX
	tx := a.db.BeginTx()

	if tx.Error != nil {
		errMsg := fmt.Sprintf("something went wrong tx event %v", tx.Error)
		return domain.Event{}, status.Errorf(codes.Internal, errMsg)
	}

	// Defer rollback in case of an error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result, err := a.db.Create(event)
	if err != nil {
		// Rollback the transaction for the database service in case of an error
		tx.Rollback()
		return result, err
	}

	// Create organiser
	err = a.organiser.CreateOrganiser(userId, result.ID)
	if err != nil {
		// Rollback the transaction for the database service in case of an error
		tx.Rollback()
		return result, err
	}

	// Commit the transaction
	tx.Commit()

	return result, nil

}

func (a *Application) CreateOrganiser(eventId uint64, userid uint64) error {

	return errors.New("not implemented")

}

func (a *Application) GetAllEvents(params domain.Params) ([]domain.Event, error) {
	events, err := a.db.GetEvents(params)

	return events, err

}
