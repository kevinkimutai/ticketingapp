package api

import "github.com/kevinkimutai/ticketingapp/auth/ports"

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) Signup() {}

func (a *Application) Login() {}
