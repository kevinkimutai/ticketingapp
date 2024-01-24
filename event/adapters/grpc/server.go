package grpc

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/charmbracelet/log"

	"github.com/kevinkimutai/ticketingapp/event/ports"
	eventproto "github.com/kevinkimutai/ticketingapp/event/proto/golang/event"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	api    ports.APIPort
	auth   ports.AuthPort
	port   int
	server *grpc.Server
	eventproto.UnimplementedEventServer
}
type contextkey string

const userIDKey contextkey = "userid"

func NewAdapter(api ports.APIPort, port int, auth ports.AuthPort) *Adapter {

	return &Adapter{api: api, port: port, auth: auth}
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(a.JWTAuthInterceptor))

	a.server = grpcServer

	eventproto.RegisterEventServer(grpcServer, a)
	if os.Getenv("ENV") == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
	fmt.Printf("GRPC server running on PORT: %v", a.port)
}

func (a *Adapter) Verify(token string) (uint64, error) {
	userId, err := a.auth.Verify(token)

	return userId, err

}

func (a *Adapter) JWTAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "metadata not provided")
	}
	//Check If User is Authenticated
	authorization, exists := md["authorization"]
	if !exists || len(authorization) == 0 {
		return nil, status.Errorf(codes.PermissionDenied, "unauthorised.login")
	}

	startsWith := "Bearer"

	if strings.HasPrefix(authorization[0], startsWith) {
		tokenStr := strings.TrimPrefix(authorization[0], startsWith)
		token := strings.TrimSpace(tokenStr)

		userID, err := a.auth.Verify(token)
		if err != nil {
			return nil, err
		}

		// Add user ID to context
		ctx = context.WithValue(ctx, userIDKey, userID)

		return handler(ctx, req)
	}
	return nil, status.Errorf(codes.PermissionDenied, "unauthorised.login")
}
