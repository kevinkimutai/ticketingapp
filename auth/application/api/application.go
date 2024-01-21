package api

import (
	"github.com/kevinkimutai/ticketingapp/auth/application/domain"
	"github.com/kevinkimutai/ticketingapp/auth/ports"
)

type Application struct {
	db  ports.DBPort
	api ports.APIPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) Signup(user domain.User) (domain.User, error) {
	result, err := a.db.CreateUser(user)

	return result, err

}

func (a *Application) Login(user domain.LoginUser) (string, error) {
	tokenStr, err := a.db.Login(user)

	return tokenStr, err
}
