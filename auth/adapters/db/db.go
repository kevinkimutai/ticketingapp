package db

import (
	"fmt"

	"github.com/kevinkimutai/ticketingapp/auth/application/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dbString string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dbString), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	err := db.AutoMigrate(&User{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}

func (a *Adapter) CreateUser(user domain.User) (domain.User, error) {
	err := a.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
