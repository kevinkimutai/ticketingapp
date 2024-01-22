package grpc

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/kevinkimutai/ticketingapp/event/ports"
	eventproto "github.com/kevinkimutai/ticketingapp/event/proto/golang/event"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	eventproto.UnimplementedEventServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
	grpcServer := grpc.NewServer()

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

//TODO:FINISH AUTHENTICATION
// func jwtAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
// ) (interface{}, error) {
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return nil, status.Errorf(codes.InvalidArgument, "metadata not provided")
// 	}
// 	//Check If User is Authenticated

// 	return handler(ctx, req)
// }
