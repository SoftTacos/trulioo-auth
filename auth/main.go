package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	gopg "github.com/go-pg/pg/v10"
	c "github.com/softtacos/trulioo-auth/auth/controller"
	d "github.com/softtacos/trulioo-auth/auth/dao"
	h "github.com/softtacos/trulioo-auth/auth/handler"
	clients "github.com/softtacos/trulioo-auth/grpc"
	v1 "github.com/softtacos/trulioo-auth/grpc/auth/v1"
	uv1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
	"github.com/softtacos/trulioo-auth/util"
	"google.golang.org/grpc"
)

const (
	usersUrlEnv = "USERS_CLIENT_ADDRESS"
	dbUrlEnv    = "DB_URL"
	grpcPortEnv = "GRPC_PORT"
)

// using this to fake the envs that would normally be set in a chart
// init gets called before anything else
func init() {
	os.Setenv("DB_URL", "postgres://postgres:postgres@localhost/tl_auth?sslmode=disable")
	os.Setenv("USERS_CLIENT_ADDRESS", ":11001")
	os.Setenv("GRPC_PORT", "11000")
}

func main() {

	var (
		usersClient uv1.UsersServiceClient
		dao         d.Dao
		controller  c.AuthController
		handler     *h.AuthHandler
	)

	karen := &ServiceManager{
		ClientsSetup: func(m *ServiceManager) (err error) {
			usersClientConn := m.CreateClientConnection(os.Getenv(usersUrlEnv))
			usersClient = uv1.NewUsersServiceClient(usersClientConn)
			return
		},
		ServiceSetup: func(m *ServiceManager) (err error) {
			controller = c.NewAuthController(dao, usersClient)
			handler = h.NewAuthHandler(controller)
			v1.RegisterAuthServiceServer(m.CreateGrpcServer(os.Getenv(grpcPortEnv)), handler)
			return
		},
		DatabaseSetup: func(m *ServiceManager) (err error) {
			db := m.CreateDbConnection(os.Getenv(dbUrlEnv))
			dao = d.NewDao(db)
			return
		},
	}

	defer karen.Stop()
	karen.Start()
}

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
	db, err = util.CreateGoPgDB(os.Getenv(dbUrlEnv))
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
