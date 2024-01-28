package grpc

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/kevinkimutai/ticketingapp/event/application/domain"
	eventproto "github.com/kevinkimutai/ticketingapp/event/proto/golang/event"
)

func (a Adapter) CreateEvent(ctx context.Context, req *eventproto.CreateEventRequest) (*eventproto.CreateEventResponse, error) {

	request := domain.Event{Name: req.Name, PosterImgURL: req.PosterImg, Venue: req.Venue, Town: req.Town, Longitude: float64(req.Longitude), Latitude: float64(req.Latitude), CreatedAt: time.Now()}

	newEvent, err := domain.NewEvent(request)
	if err != nil {
		return nil, err
	}

	userId := ctx.Value("userid").(uint64)
	log.Info(userId)

	result, err := a.api.CreateEvent(userId, newEvent)
	if err != nil {
		return nil, err
	}

	return &eventproto.CreateEventResponse{EventId: result.ID}, nil

}

//CREATETICKET

//CREATETICKETTYPE
