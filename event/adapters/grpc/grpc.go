package grpc

import (
	"context"

	eventproto "github.com/kevinkimutai/ticketingapp/event/proto/golang/event"
)

func (a Adapter) CreateEvent(ctx context.Context, req *eventproto.CreateEventRequest) (*eventproto.CreateEventResponse, error) {

}
