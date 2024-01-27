package organiser

import (
	"errors"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"

	organiserproto "github.com/kevinkimutai/ticketingapp/event/proto/golang/organiser"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	organiser organiserproto.OrganiserClient
}

func NewAdapter(organiserServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
		grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
		grpc_retry.WithMax(5),
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
	)))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(organiserServiceUrl, opts...)
	if err != nil {
		return nil, err
	}

	client := organiserproto.NewOrganiserClient(conn)
	return &Adapter{organiser: client}, nil
}

func (adapter *Adapter) CreateOrganiser(eventId uint64, userid uint64) error {
	//OrganiserRequest

	return errors.New("abcdef")
}
