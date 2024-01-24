package ports

import "github.com/kevinkimutai/ticketingapp/auth/application/domain"

type DBPort interface {
	CreateUser(user domain.User) (domain.User, error)
	Login(user domain.LoginUser) (string, error)
	VerifyJWT(token string) (domain.User, error)
}
