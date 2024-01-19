package ports

import "github.com/kevinkimutai/ticketingapp/auth/application/domain"

type APIPort interface {
	Signup(user domain.User) (domain.User, error)
	Login()
}
