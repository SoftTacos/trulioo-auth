package main

import (
	"os"

	c "github.com/softtacos/trulioo-auth/auth/controller"
	d "github.com/softtacos/trulioo-auth/auth/dao"
	h "github.com/softtacos/trulioo-auth/auth/handler"
	v1 "github.com/softtacos/trulioo-auth/grpc/auth/v1"
	uv1 "github.com/softtacos/trulioo-auth/grpc/users/v1"
	sm "github.com/softtacos/trulioo-auth/util/service-manager"
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

	karen := &sm.ServiceManager{
		ClientsSetup: func(m *sm.ServiceManager) (err error) {
			usersClientConn := m.CreateClientConnection(os.Getenv(usersUrlEnv))
			usersClient = uv1.NewUsersServiceClient(usersClientConn)
			return
		},
		ServiceSetup: func(m *sm.ServiceManager) (err error) {
			controller = c.NewAuthController(dao, usersClient)
			handler = h.NewAuthHandler(controller)
			v1.RegisterAuthServiceServer(m.CreateGrpcServer(os.Getenv(grpcPortEnv)), handler)
			return
		},
		DatabaseSetup: func(m *sm.ServiceManager) (err error) {
			db := m.CreateDbConnection(os.Getenv(dbUrlEnv))
			dao = d.NewDao(db)
			return
		},
	}

	defer karen.Stop()
	karen.Start()
}
