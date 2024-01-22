package grpc

import (
	"context"
	"time"

	"github.com/kevinkimutai/ticketingapp/auth/application/domain"
	authproto "github.com/kevinkimutai/ticketingapp/auth/proto/golang"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (a Adapter) Login(ctx context.Context, req *authproto.LoginRequest) (*authproto.LoginResponse, error) {

	request := domain.LoginUser{Email: req.Email, Password: req.Password}

	newLogin, err := domain.NewLogin(request)
	if err != nil {
		return nil, err
	}

	//Check Email Is Valid
	err = domain.CheckEmail(newLogin.Email)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "email address is invalid")
	}

	//Check If Email Exists In DB
	tokenStr, err := a.api.Login(newLogin)
	if err != nil {
		return nil, err
	}
	return &authproto.LoginResponse{Token: tokenStr}, nil
}

func (a Adapter) Signup(ctx context.Context, req *authproto.SignUpRequest) (*authproto.SignUpResponse, error) {

	request := domain.User{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Password: req.Password, CreatedAt: time.Now()}

	newUser, err := domain.NewUser(request)
	if err != nil {
		return nil, err
	}

	//Check Email Is Valid
	err = domain.CheckEmail(newUser.Email)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid email address")
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
