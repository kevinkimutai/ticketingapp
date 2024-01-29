package gateway

import (
	"context"
	"fmt"

	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	eventproto "github.com/kevinkimutai/ticketingapp/event/proto/golang/event"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	port int
	grpc int
}

func NewAdapter(grpc int, port int) *Adapter {
	return &Adapter{grpc: grpc, port: port}
}

func (a Adapter) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	grpcAddr := fmt.Sprintf("localhost:%d", a.grpc)

	err := eventproto.RegisterEventHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	portAddr := fmt.Sprintf(":%d", a.port)

	err = http.ListenAndServe(portAddr, mux)
	if err != nil {
		return err
	}

	return err

}
