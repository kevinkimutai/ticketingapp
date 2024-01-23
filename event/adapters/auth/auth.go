package auth

import (
	"context"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	authproto "github.com/kevinkimutai/ticketingapp/event/proto/golang/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	auth authproto.AuthProtoClient
}

func NewAdapter(authServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
	)))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(authServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := authproto.NewAuthProtoClient(conn)
	return &Adapter{auth: client}, nil
}

func (a *Adapter) Verify(token string) error {
	_, err := a.auth.VerifyJWT(context.Background(),
		&authproto.VerifyTokenRequest{
			Token: token,
		})

	return err
}
