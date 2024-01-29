package domain

import (
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Event struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	PosterImgURL string `json:"poster_img"`

	Venue     string  `json:"venue"`
	Town      string  `json:"town"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	CreatedAt time.Time
}

type Params struct {
	PageNumber int64  `json:"page_number"`
	PageSize   int64  `json:"page_size"`
	Search     string `json:"search"`
}

func NewEvent(e Event) (Event, error) {
	if e.Name == "" || e.PosterImgURL == "" || e.Town == "" || e.Venue == "" {
		return e, status.Errorf(codes.InvalidArgument, "missing input fields in event")
	}

	return e, nil
}

func NewParams(p Params) Params {

	if p.PageNumber == 0 {
		p.PageNumber = 1
	}

	if p.PageSize == 0 {
		p.PageSize = 20
	}

	if p.PageSize > 100 {
		p.PageSize = 100
	}

	return p
}
