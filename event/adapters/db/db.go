package db

import (
	"fmt"

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
	PosterImgURL string   `json:"poster_img"`
	ImagesURL    []string `json:"images_url"`
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
