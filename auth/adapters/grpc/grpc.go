package grpc

import (
	"context"
	"errors"
	"time"

	"github.com/kevinkimutai/ticketingapp/auth/application/domain"
	authproto "github.com/kevinkimutai/ticketingapp/auth/proto/golang"
)

func (a Adapter) Login(ctx context.Context, req *authproto.LoginRequest) (*authproto.LoginResponse, error) {

	return nil, nil
}

func (a Adapter) Signup(ctx context.Context, req *authproto.SignUpRequest) (*authproto.SignUpResponse, error) {

	request := domain.User{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Password: req.Password, CreatedAt: time.Now()}

	newUser, err := domain.NewUser(request)
	if err != nil {
		return nil, err
	}

	//Check Email Is Valid
	err = newUser.CheckEmail()
	if err != nil {
		return nil, errors.New("invalid email address ")
	}

	//Check Password Strength
	err = newUser.CheckPasswordStrength()
	if err != nil {
		return nil, err
	}

	//Hash Password
	hashedUser, err := newUser.HashPassword()
	if err != nil {
		return nil, err

	}

	result, err := a.api.Signup(hashedUser)
	if err != nil {
		return nil, err

	}
	return &authproto.SignUpResponse{UserId: result.ID}, nil

}
