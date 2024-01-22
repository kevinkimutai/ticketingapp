package domain

import (
	"errors"
	"time"
)

type Event struct {
	ID           int64    `json:"id"`
	Name         string   `json:"name"`
	PosterImgURL string   `json:"poster_img"`
	ImagesURL    []string `json:"images_url"`
	Venue        string   `json:"venue"`
	Town         string   `json:"town"`
	Longitude    float64  `json:"longitude"`
	Latitude     float64  `json:"latitude"`
	CreatedAt    time.Time
}

func NewEvent(e Event) (Event, error) {
	if e.Name == "" || e.PosterImgURL == "" || e.Town == "" || e.Venue == "" {
		return e, errors.New("missing fields in event")
	}

	return event, nil
}
