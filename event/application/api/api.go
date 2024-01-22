package api

import "github.com/kevinkimutai/ticketingapp/event/ports"

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
