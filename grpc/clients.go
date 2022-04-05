package grpc

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/newrelic/go-agent/v3/integrations/nrgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	maxGrpcServerRecvSize = 20 * 1024 * 1024
	maxGrpcServerSendSize = 20 * 1024 * 1024
	withPerRetryTimeout   = 1 * time.Minute
)

type ClientOptions struct {
	MaxGrpcServerRecvSize int
	MaxGrpcServerSendSize int
	WithPerRetryTimeout   time.Duration
}

type ClientManager struct {
	connections *[]*grpc.ClientConn
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		connections: &[]*grpc.ClientConn{},
	}
}

// Create creates a connection for a gRPC client
func (f *ClientManager) Create(url string, opts ...ClientOptions) *grpc.ClientConn {
	connection, err := createRpcConnection(url, opts...)
	if err != nil {
		log.Panic(err)
	}
	*f.connections = append(*f.connections, connection)
	return connection
}

func createRpcConnection(target string, opts ...ClientOptions) (connection *grpc.ClientConn, err error) {
	log.Println("Creating RPC Client for " + target)

	opt := ClientOptions{
		MaxGrpcServerRecvSize: maxGrpcServerRecvSize,
		MaxGrpcServerSendSize: maxGrpcServerSendSize,
		WithPerRetryTimeout:   withPerRetryTimeout,
	}
	if len(opts) > 0 {
		if opts[0].MaxGrpcServerRecvSize > 0 {
			opt.MaxGrpcServerRecvSize = opts[0].MaxGrpcServerRecvSize
		}
		if opts[0].MaxGrpcServerSendSize > 0 {
			opt.MaxGrpcServerSendSize = opts[0].MaxGrpcServerSendSize
		}
		if opts[0].WithPerRetryTimeout > 0 {
			opt.WithPerRetryTimeout = opts[0].WithPerRetryTimeout
		}
	}

	if target == "" || !strings.Contains(target, ":") {
		log.Panic(fmt.Sprintf("invalid url given for gRPC client: %s", target))
	}

	retryOptions := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)),
		grpc_retry.WithCodes(codes.Aborted, codes.Canceled, codes.Unavailable),
		grpc_retry.WithPerRetryTimeout(opt.WithPerRetryTimeout),
	}

	connection, err = grpc.Dial(
		target,
		// IRL we would want to use secure gRPC OR have the REST authentication be https.
		// I'm assuming that the request has gone through some sort of gateway into the cluster and is safe
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(opt.MaxGrpcServerRecvSize),
			grpc.MaxCallSendMsgSize(opt.MaxGrpcServerSendSize)),
		grpc.WithChainUnaryInterceptor(
			nrgrpc.UnaryClientInterceptor,
			grpc_retry.UnaryClientInterceptor(retryOptions...)),
	/*grpc.WithChainStreamInterceptor(
	)*/
	)

	if err != nil {
		log.Println("could not connect to gRPC client: ", err)
	}
	return
}

func validateUrl(url string) (err error) {
	if !strings.Contains(url, ":") {
		return errors.New("URL does not have a port")
	}

	return
}
