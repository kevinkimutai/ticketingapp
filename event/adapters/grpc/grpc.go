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

func (a Adapter) GetEvents(ctx context.Context, req *eventproto.GetEventsRequest) (*eventproto.GetEventsResponse, error) {
	request := domain.Params{PageNumber: req.PageNumber, PageSize: req.PageSize}

	params := domain.NewParams(request)

	events, err := a.api.GetAllEvents(params)
	if err != nil {
		return nil, err
	}

	response := &eventproto.GetEventsResponse{
		Event: []*eventproto.EventType{},
	}

	//Convert domain.Event to eventProto.GetEventResponse
	for _, event := range events {
		protoEvent := ConvertDomainEventToProtoResponse(event)
		response.Event = append(response.Event, protoEvent)
	}

	return &eventproto.GetEventsResponse{Event: response.Event}, nil

}

//TODO:CHECK ON TIME FORMAT

func ConvertDomainEventToProtoResponse(event domain.Event) *eventproto.EventType {
	protoEvent := &eventproto.EventType{
		ID:        event.ID,
		Name:      event.Name,
		Venue:     event.Venue,
		Town:      event.Town,
		Longitude: float32(event.Latitude),
		Latitude:  float32(event.Latitude),
		PosterImg: event.PosterImgURL,
		StartTime: event.StartTime.String(),
		EndTime:   event.EndTime.String(),
	}
	return protoEvent
}
