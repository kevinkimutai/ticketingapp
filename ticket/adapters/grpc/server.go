package grpc

import (
	"fmt"
	"net"
	"os"

	"github.com/charmbracelet/log"
	"github.com/kevinkimutai/ticketingapp/ticket/ports"
	ticketproto "github.com/kevinkimutai/ticketingapp/ticket/proto/golang"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	ticketproto.UnimplementedTicketProtoServer
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

	ticketproto.RegisterTicketProtoServer(grpcServer, a)
	if os.Getenv("ENV") == "development" {
		reflection.Register(grpcServer)
	}
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port ")
	}
	log.Info("GRPC server running on PORT: %v", a.port)
}
