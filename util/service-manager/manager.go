package service_manager

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	gopg "github.com/go-pg/pg/v10"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcServer struct {
	server   *grpc.Server
	listener net.Listener
}

func (s *grpcServer) Serve() {
	go s.server.Serve(s.listener)
}

func (s *grpcServer) Stop() {
	s.server.GracefulStop()
	s.listener.Close()
}

type ServiceManager struct {
	ServiceSetup  func(m *ServiceManager)
	ClientsSetup  func(m *ServiceManager)
	DatabaseSetup func(m *ServiceManager)

	servers           []*grpcServer
	dbConnections     []*gopg.DB
	clientConnections []*grpc.ClientConn
}

func (m *ServiceManager) CreateGrpcServer(port string) (server *grpc.Server) {
	grpcListener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server = grpc.NewServer()
	m.servers = append(m.servers, &grpcServer{
		server:   server,
		listener: grpcListener,
	})
	return
}

func (m *ServiceManager) CreateDbConnection(url string) (db *gopg.DB) {
	var err error
	db, err = createGoPgDB(url)
	if err != nil {
		log.Panic("error creating DB: ", err)
		return
	}
	m.dbConnections = append(m.dbConnections, db)
	return
}

func (m *ServiceManager) CreateClientConnection(url string) (connection *grpc.ClientConn) {
	var err error
	connection, err = createRpcConnection(url)
	if err != nil {
		log.Panic("failed to connect to client: ", err.Error())
		return
	}
	m.clientConnections = append(m.clientConnections)

	return
}

func (m *ServiceManager) Start() {
	if m.ClientsSetup != nil {
		m.ClientsSetup(m)
	}

	if m.DatabaseSetup != nil {
		m.DatabaseSetup(m)
	}

	if m.ServiceSetup != nil {
		m.ServiceSetup(m)
	}
	for _, server := range m.servers {
		server.Serve()
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	log.Println("SETUP COMPLETE")
	<-sigs
	log.Println("SHUTTING DOWN")
}

func (m *ServiceManager) Stop() {
	var err error

	for _, dbConn := range m.dbConnections {
		err = dbConn.Close()
		if err != nil {
			log.Println("failed to close db connection: ", err.Error())
		}
	}

	for _, server := range m.servers {
		server.Stop()
	}
}

func createRpcConnection(target string) (connection *grpc.ClientConn, err error) {
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
