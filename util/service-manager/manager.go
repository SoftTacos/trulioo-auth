package service_manager

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	gopg "github.com/go-pg/pg/v10"
	"google.golang.org/grpc"
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
	ServiceSetup  func(m *ServiceManager) error
	ClientsSetup  func(m *ServiceManager) error
	DatabaseSetup func(m *ServiceManager) error

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
	db, err = CreateGoPgDB(os.Getenv(dbUrlEnv))
	if err != nil {
		log.Panic("error creating DB: ", err)
		return
	}
	m.dbConnections = append(m.dbConnections, db)
	return
}

func (m *ServiceManager) CreateClientConnection(url string) (connection *grpc.ClientConn) {
	var err error
	connection, err = clients.CreateRpcConnection(url)
	if err != nil {
		log.Panic("failed to connect to client: ", err.Error())
		return
	}
	m.clientConnections = append(m.clientConnections)

	return
}

func (m *ServiceManager) Start() {
	var err error

	err = m.ClientsSetup(m)
	if err != nil {
		log.Panic("failed to setup clients: ", err.Error())
	}

	err = m.DatabaseSetup(m)
	if err != nil {
		log.Panic("failed to setup databases: ", err.Error())
	}

	err = m.ServiceSetup(m)
	if err != nil {
		log.Panic("failed to setup service: ", err.Error())
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
