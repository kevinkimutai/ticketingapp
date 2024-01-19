package grpc

import (
	"context"
	"fmt"

	authproto "github.com/kevinkimutai/ticketingapp/auth/proto/golang"
)

func (a Adapter) Login(ctx context.Context, req *authproto.LoginRequest) (*authproto.LoginResponse, error) {
	fmt.Println("REQ", req)

	return nil, nil
}
