package grpc

import (
	"errors"
	"log"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// creates basic standard grpc connection
func CreateRpcConnection(target string) (connection *grpc.ClientConn, err error) {
	log.Println("Creating RPC Client for " + target)
	if err = validateUrl(target); err != nil {
		return
	}

	connection, err = grpc.Dial(
		target,
		// IRL we would want to use secure gRPC OR have the REST authentication be https.
		// I'm assuming that the request has gone through some sort of gateway into the cluster and is safe
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(),
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
